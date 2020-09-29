package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dghubble/gologin"
	"github.com/dghubble/gologin/github"
	oauth2Login "github.com/dghubble/gologin/oauth2"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
	githuboauth2 "golang.org/x/oauth2/github"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server/models"
)

const (
	methodGET  = "GET"
	methodPOST = "POST"
	methodPUT  = "PUT"
)

const (
	sessionName      = "atlas_session"
	sessionUserKey   = "github_id"
	sessionUserLogin = "github_login"
)

// Service defines the encapsulating Atlas service. It wraps a router which is
// responsible for handling API requests with a given controller that interacts
// with Atlas models. The Service is responsible for establishing a database
// connection and managing session cookies.
type Service struct {
	logger       zerolog.Logger
	cfg          config.Config
	db           *gorm.DB
	cookieCfg    gologin.CookieConfig
	sessionStore *sessions.CookieStore
	oauth2Cfg    *oauth2.Config
	validate     *validator.Validate
	router       *mux.Router
	server       *http.Server
}

func NewService(logger zerolog.Logger, cfg config.Config) (*Service, error) {
	dbLogger := NewDBLogger(logger).LogMode(gormlogger.Silent)

	cookieCfg := gologin.DefaultCookieConfig
	sessionStore := sessions.NewCookieStore([]byte(cfg.String(config.FlagSessionKey)), nil)
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = true
	sessionStore.Options.MaxAge = 3600 * 24 * 7 // 1 week

	if cfg.Bool(config.FlagDev) {
		dbLogger = dbLogger.LogMode(gormlogger.Info)
		cookieCfg = gologin.DebugOnlyCookieConfig
		sessionStore.Options.Secure = false
	}

	db, err := gorm.Open(postgres.Open(cfg.String(config.FlagDatabaseURL)), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, err
	}

	service := &Service{
		logger:       logger.With().Str("module", "server").Logger(),
		cfg:          cfg,
		db:           db,
		cookieCfg:    cookieCfg,
		sessionStore: sessionStore,
		validate:     validator.New(),
		router:       mux.NewRouter(),
		oauth2Cfg: &oauth2.Config{
			ClientID:     cfg.String(config.FlagGHClientID),
			ClientSecret: cfg.String(config.FlagGHClientSecret),
			RedirectURL:  cfg.String(config.FlagGHRedirectURL),
			Endpoint:     githuboauth2.Endpoint,
		},
	}

	service.registerV1Routes()
	return service, nil
}

// Start starts the atlas service as a blocking process.
func (s *Service) Start() error {
	s.server = &http.Server{
		Handler:      s.router,
		Addr:         s.cfg.String(config.FlagListenAddr),
		WriteTimeout: s.cfg.Duration(config.FlagHTTPReadTimeout),
		ReadTimeout:  s.cfg.Duration(config.FlagHTTPWriteTimeout),
	}

	s.logger.Info().Str("address", s.server.Addr).Msg("starting atlas server...")
	return s.server.ListenAndServe()
}

// Cleanup performs server cleanup. If the internal HTTP server is non-nil, the
// server will be shutdown after a grace period deadline.
func (s *Service) Cleanup() {
	if s.server != nil {
		// create a deadline to wait for all existing requests to finish
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Do not block if no connections exist, but otherwise, we will wait until
		// the timeout deadline.
		_ = s.server.Shutdown(ctx)
	}
}

func (s *Service) registerV1Routes() {
	// Create a versioned sub-router. All API routes will be mounted under this
	// sub-router.
	v1Router := s.router.PathPrefix("/api/v1").Subrouter()

	// build middleware chain
	mChain := buildMiddleware(s.logger)

	// unauthenticated routes
	v1Router.Handle(
		"/modules/search",
		mChain.ThenFunc(s.SearchModules()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}", "q", "{q}").Methods(methodGET)

	v1Router.Handle(
		"/modules",
		mChain.ThenFunc(s.GetAllModules()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(methodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}",
		mChain.ThenFunc(s.GetModuleByID()),
	).Methods(methodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/versions",
		mChain.ThenFunc(s.GetModuleVersions()),
	).Methods(methodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/authors",
		mChain.ThenFunc(s.GetModuleAuthors()),
	).Methods(methodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/keywords",
		mChain.ThenFunc(s.GetModuleKeywords()),
	).Methods(methodGET)

	v1Router.Handle(
		"/users/{id:[0-9]+}",
		mChain.ThenFunc(s.GetUserByID()),
	).Methods(methodGET)

	v1Router.Handle(
		"/users/{id:[0-9]+}/modules",
		mChain.ThenFunc(s.GetUserModules()),
	).Methods(methodGET)

	v1Router.Handle(
		"/users",
		mChain.ThenFunc(s.GetAllUsers()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(methodGET)

	v1Router.Handle(
		"/keywords",
		mChain.ThenFunc(s.GetAllKeywords()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(methodGET)

	// authenticated routes
	v1Router.Handle(
		"/modules",
		mChain.ThenFunc(s.UpsertModule()),
	).Methods(methodPUT)

	// session routes
	v1Router.Handle(
		"/session/start",
		mChain.Then(s.BeginSession()),
	).Methods(methodGET)

	v1Router.Handle(
		"/session/authorize",
		mChain.Then(s.AuthorizeSession()),
	).Methods(methodGET)

	v1Router.Handle(
		"/session/logout",
		mChain.ThenFunc(s.LogoutSession()),
	).Methods(methodPOST)
}

// UpsertModule implements a request handler to create or update a Cosmos SDK
// module.
func (s *Service) UpsertModule() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ghUserID, ok, err := s.authorize(req)
		if err != nil || !ok {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		var request ModuleRequest
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("failed to read request: %w", err))
			return
		}

		if err := s.validate.Struct(request); err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid request: %w", transformValidationError(err)))
			return
		}

		module := ModuleFromRequest(request)
		publisher, err := models.User{GithubUserID: sql.NullInt64{Int64: ghUserID, Valid: true}}.Query(s.db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get publisher: %w", err))
			return
		}

		// The publisher must already be an existing owner or must have accepted an
		// invitation by an existing owner.
		//
		// TODO: Handle invitations to allow other users to update existing modules.
		record, err := models.Module{Name: module.Name, Team: module.Team}.Query(s.db)
		if err == nil {
			// the module already exists so we check if the publisher is an owner
			var isOwner bool
			for i := 0; i < len(record.Owners) && !isOwner; i++ {
				if record.Owners[i].GithubUserID == publisher.GithubUserID {
					isOwner = true
				}
			}

			if !isOwner {
				respondWithError(w, http.StatusBadRequest, errors.New("publisher must be an owner of the module"))
				return
			}

			module.Owners = record.Owners
		} else {
			module.Owners = []models.User{publisher}
		}

		module, err = module.Upsert(s.db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module)
	}
}

// GetModuleByID implements a request handler to retrieve a module by ID.
func (s *Service) GetModuleByID() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(s.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			respondWithError(w, code, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module)
	}
}

// SearchModules implements a request handler to retrieve a set of module objects
// by search criteria.
func (s *Service) SearchModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := parsePagination(req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		query := req.URL.Query().Get("q")

		modules, err := models.SearchModules(s.db, query, cursor, limit)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := NewPaginationResponse(len(modules), limit, cursor, modules)
		respondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetAllModules implements a request handler returning a paginated set of
// modules.
func (s *Service) GetAllModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := parsePagination(req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		modules, err := models.GetAllModules(s.db, cursor, limit)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := NewPaginationResponse(len(modules), limit, cursor, modules)
		respondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetModuleVersions implements a request handler to retreive a module's set of
// versions by ID.
func (s *Service) GetModuleVersions() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(s.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module.Versions)
	}
}

// GetModuleAuthors implements a request handler to retreive a module's set of
// authors by ID.
func (s *Service) GetModuleAuthors() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(s.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module.Authors)
	}
}

// GetModuleKeywords implements a request handler to retreive a module's set of
// keywords by ID.
func (s *Service) GetModuleKeywords() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(s.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module.Keywords)
	}
}

// GetUserByID implements a request handler to retrieve a user by ID.
func (s *Service) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %w", err))
			return
		}

		user, err := models.GetUserByID(s.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, user)
	}
}

// GetAllUsers implements a request handler returning a paginated set of
// users.
func (s *Service) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := parsePagination(req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		users, err := models.GetAllUsers(s.db, cursor, limit)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := NewPaginationResponse(len(users), limit, cursor, users)
		respondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetUserModules implements a request handler to retrieve a set of modules
// authored by a given user by ID.
func (s *Service) GetUserModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %w", err))
			return
		}

		modules, err := models.GetUserModules(s.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, modules)
	}
}

// GetAllKeywords implements a request handler returning a paginated set of
// keywords.
func (s *Service) GetAllKeywords() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := parsePagination(req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		keywords, err := models.GetAllKeywords(s.db, cursor, limit)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := NewPaginationResponse(len(keywords), limit, cursor, keywords)
		respondWithJSON(w, http.StatusOK, paginated)
	}
}

// BeginSession returns a request handler to begin a user session via Github
// OAuth authentication. The user must either grant or reject access. Upon
// granting access, Github will perform a callback where we create a session
// and obtain a token.
func (s *Service) BeginSession() http.Handler {
	return github.StateHandler(s.cookieCfg, github.LoginHandler(s.oauth2Cfg, nil))
}

// AuthorizeSession returns a callback request handler for Github OAuth user
// authentication. After a user grants access, this callback handler will be
// executed. A session cookie will be saved and sent to the client. A user record
// will also be upserted.
func (s *Service) AuthorizeSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		token, err := oauth2Login.TokenFromContext(ctx)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get github token: %w", err))
			return
		}

		githubUser, err := github.UserFromContext(ctx)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get github user: %w", err))
			return
		}

		user := models.User{
			Name:              githubUser.GetLogin(),
			GithubUserID:      sql.NullInt64{Int64: githubUser.GetID(), Valid: true},
			GravatarID:        githubUser.GetGravatarID(),
			AvatarURL:         githubUser.GetAvatarURL(),
			GithubAccessToken: sql.NullString{String: token.AccessToken, Valid: true},
		}
		if _, err := user.Upsert(s.db); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to upsert user: %w", err))
			return
		}

		session, err := s.sessionStore.Get(req, sessionName)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
			return
		}

		session.Values[sessionUserKey] = githubUser.GetID()
		session.Values[sessionUserLogin] = githubUser.GetLogin()

		if err = session.Save(req, w); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to save session: %w", err))
			return
		}

		http.Redirect(w, req, "/", http.StatusFound)
	}

	return github.StateHandler(s.cookieCfg, github.CallbackHandler(s.oauth2Cfg, http.HandlerFunc(fn), nil))
}

// LogoutSession implements a request handler to terminate and logout of an
// existing session.
func (s *Service) LogoutSession() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := s.sessionStore.Get(req, sessionName)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
			return
		}

		// Remove session keys and set max age to -1 to trigger the deletion of the
		// cookie.
		delete(session.Values, sessionUserKey)
		delete(session.Values, sessionUserLogin)
		session.Options.MaxAge = -1
		if err = session.Save(req, w); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to save session: %w", err))
			return
		}

		http.Redirect(w, req, "/", http.StatusFound)
	}
}

// authorize attempts to authorize the given request against the session cookie
// store. If the session does not exist or if the session has been deleted, we
// treat the request as unauthorized.
func (s *Service) authorize(req *http.Request) (int64, bool, error) {
	session, err := s.sessionStore.Get(req, sessionName)
	if err != nil {
		return 0, false, fmt.Errorf("failed to get session: %w", err)
	}

	v, ok := session.Values[sessionUserKey]
	if !ok {
		return 0, false, errors.New("unauthorized")
	}

	return v.(int64), true, nil
}
