package server

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/InVisionApp/go-health/v2"
	"github.com/InVisionApp/go-health/v2/handlers"
	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/github"
	oauth2login "github.com/dghubble/gologin/v2/oauth2"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"
	githuboauth2 "golang.org/x/oauth2/github"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server/models"
)

const (
	methodGET    = "GET"
	methodPOST   = "POST"
	methodPUT    = "PUT"
	methodDELETE = "DELETE"

	sessionName     = "atlas_session"
	sessionGithubID = "github_id"
	sessionUserID   = "user_Id"

	bearerSchema = "Bearer "
)

var (
	// MaxTokens defines the maximum number of API tokens a user can create.
	MaxTokens int64 = 100
)

// Service defines the encapsulating Atlas service. It wraps a router which is
// responsible for handling API requests with a given controller that interacts
// with Atlas models. The Service is responsible for establishing a database
// connection and managing session cookies.
type Service struct {
	logger        zerolog.Logger
	cfg           config.Config
	db            *gorm.DB
	cookieCfg     gologin.CookieConfig
	sessionStore  *sessions.CookieStore
	oauth2Cfg     *oauth2.Config
	validate      *validator.Validate
	router        *mux.Router
	healthChecker *health.Health
	server        *http.Server
}

func NewService(logger zerolog.Logger, cfg config.Config) (*Service, error) {
	dbLogger := NewDBLogger(logger).LogMode(gormlogger.Silent)

	sessionKey, err := base64.StdEncoding.DecodeString(cfg.String(config.FlagSessionKey))
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode session key: %w", err)
	}

	cookieCfg := gologin.DefaultCookieConfig
	sessionStore := sessions.NewCookieStore(sessionKey, nil)
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

	sqlDB, _ := db.DB()
	healthChecker, err := CreateHealthChecker(sqlDB, true)
	if err != nil {
		return nil, err
	}

	service := &Service{
		logger:        logger.With().Str("module", "server").Logger(),
		cfg:           cfg,
		db:            db,
		cookieCfg:     cookieCfg,
		sessionStore:  sessionStore,
		validate:      validator.New(),
		router:        mux.NewRouter(),
		healthChecker: healthChecker,
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

	v1Router.Handle(
		"/health",
		mChain.Then(handlers.NewJSONHandlerFunc(s.healthChecker, nil)),
	).Methods(methodGET)

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

	v1Router.Handle(
		"/user/tokens",
		mChain.ThenFunc(s.CreateUserToken()),
	).Methods(methodPUT)

	v1Router.Handle(
		"/user/tokens",
		mChain.ThenFunc(s.GetUserTokens()),
	).Methods(methodGET)

	v1Router.Handle(
		"/user/tokens/{id:[0-9]+}",
		mChain.ThenFunc(s.RevokeUserToken()),
	).Methods(methodDELETE)

	// session routes
	v1Router.Handle(
		"/session/start",
		mChain.Then(s.StartSession()),
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
		authUser, ok, err := s.authorize(req)
		if err != nil || !ok {
			respondWithError(w, http.StatusUnauthorized, err)
			return
		}

		var request Manifest
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("failed to read request: %w", err))
			return
		}

		if err := s.validate.Struct(request); err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid request: %w", TransformValidationError(err)))
			return
		}

		module := ModuleFromManifest(request)

		// The publisher must already be an existing owner or must have accepted an
		// invitation by an existing owner.
		//
		// TODO: Handle invitations to allow other users to update existing modules.
		record, err := models.QueryModule(s.db, map[string]interface{}{"name": module.Name, "team": module.Team})
		if err == nil {
			// the module already exists so we check if the publisher is an owner
			var isOwner bool
			for i := 0; i < len(record.Owners) && !isOwner; i++ {
				if record.Owners[i].ID == authUser.ID {
					isOwner = true
				}
			}

			if !isOwner {
				respondWithError(w, http.StatusBadRequest, errors.New("publisher must be an owner of the module"))
				return
			}

			module.Owners = record.Owners
		} else {
			module.Owners = []models.User{authUser}
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

// GetModuleVersions implements a request handler to retrieve a module's set of
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
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			respondWithError(w, code, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module.Versions)
	}
}

// GetModuleAuthors implements a request handler to retrieve a module's set of
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
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			respondWithError(w, code, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module.Authors)
	}
}

// GetModuleKeywords implements a request handler to retrieve a module's set of
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
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			respondWithError(w, code, err)
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
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			respondWithError(w, code, err)
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
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			respondWithError(w, code, err)
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

// StartSession returns a request handler to begin a user session via Github
// OAuth authentication. The user must either grant or reject access. Upon
// granting access, Github will perform a callback where we create a session
// and obtain a token.
func (s *Service) StartSession() http.Handler {
	return github.StateHandler(s.cookieCfg, github.LoginHandler(s.oauth2Cfg, nil))
}

// AuthorizeSession returns a callback request handler for Github OAuth user
// authentication. After a user grants access, this callback handler will be
// executed. A session cookie will be saved and sent to the client. A user record
// will also be upserted.
func (s *Service) AuthorizeSession() http.Handler {
	return github.StateHandler(s.cookieCfg, github.CallbackHandler(s.oauth2Cfg, s.authorizeHandler(), nil))
}

func (s *Service) authorizeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		token, err := oauth2login.TokenFromContext(ctx)
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

		record, err := user.Upsert(s.db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to upsert user: %w", err))
			return
		}

		session, err := s.sessionStore.Get(req, sessionName)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
			return
		}

		session.Values[sessionGithubID] = githubUser.GetID()
		session.Values[sessionUserID] = record.ID

		if err = session.Save(req, w); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to save session: %w", err))
			return
		}

		http.Redirect(w, req, "/", http.StatusFound)
	}
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
		delete(session.Values, sessionGithubID)
		delete(session.Values, sessionUserID)
		session.Options.MaxAge = -1

		if err = session.Save(req, w); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to save session: %w", err))
			return
		}

		http.Redirect(w, req, "/", http.StatusFound)
	}
}

// CreateUserToken implements a request handler that creates a new API token for
// the authenticated user.
func (s *Service) CreateUserToken() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := s.authorize(req)
		if err != nil || !ok {
			respondWithError(w, http.StatusUnauthorized, err)
			return
		}

		numTokens := authUser.CountTokens(s.db)
		if numTokens >= MaxTokens {
			respondWithError(w, http.StatusBadRequest, errors.New("maximum number of user API tokens reached"))
			return
		}

		token, err := authUser.CreateToken(s.db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, token)
	}
}

// GetUserTokens implements a request handler returning all of an authenticated
// user's tokens.
func (s *Service) GetUserTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := s.authorize(req)
		if err != nil || !ok {
			respondWithError(w, http.StatusUnauthorized, err)
			return
		}

		tokens, err := authUser.GetTokens(s.db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, tokens)
	}
}

// RevokeUserToken implements a request handler revoking a specific token from
// the authorized user.
func (s *Service) RevokeUserToken() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := s.authorize(req)
		if err != nil || !ok {
			respondWithError(w, http.StatusUnauthorized, err)
			return
		}

		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		token, err := models.QueryUserToken(s.db, map[string]interface{}{"id": id, "user_id": authUser.ID, "revoked": false})
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			respondWithError(w, code, err)
			return
		}

		token, err = token.Revoke(s.db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, token)
	}
}

// authorize attempts to authorize the given request against the session cookie
// store or a bearer authorization header. If the session cookie does not exist,
// or the session has been deleted, or the supplied bearer authorization header
// is invalid, we treat the request as unauthorized and return false. Otherwise,
// we return the user record ID and true with no error to indicate successful
// authorization.
func (s *Service) authorize(req *http.Request) (models.User, bool, error) {
	session, err := s.sessionStore.Get(req, sessionName)
	if err != nil {
		return models.User{}, false, fmt.Errorf("failed to get session: %w", err)
	}

	var userID uint

	// check for a valid session cookie or bearer authorization header
	if v, ok := session.Values[sessionUserID]; ok {
		userID = v.(uint)
	} else if h := req.Header.Get("Authorization"); strings.HasPrefix(h, bearerSchema) {
		tokenStr := h[len(bearerSchema):]

		tokenUUID, err := uuid.FromString(tokenStr)
		if err != nil {
			return models.User{}, false, fmt.Errorf("failed to get parse token: %w", err)
		}

		token, err := models.QueryUserToken(s.db, map[string]interface{}{"token": tokenUUID.String(), "revoked": false})
		if err != nil {
			return models.User{}, false, err
		}

		token, err = token.IncrCount(s.db)
		if err != nil {
			return models.User{}, false, err
		}

		userID = token.UserID
	} else {
		return models.User{}, false, errors.New("unauthorized")
	}

	user, err := models.GetUserByID(s.db, userID)
	if err != nil {
		return models.User{}, false, err
	}

	return user, true, nil
}
