package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	"gorm.io/gorm"

	"github.com/cosmos/atlas/server/httputil"
	"github.com/cosmos/atlas/server/middleware"
	"github.com/cosmos/atlas/server/models"
)

const (
	sessionName     = "atlas_session"
	sessionGithubID = "github_id"
	sessionUserID   = "user_Id"

	V1APIPathPrefix = "/api/v1"
)

var (
	// MaxTokens defines the maximum number of API tokens a user can create.
	MaxTokens int64 = 100
)

// Router implements a versioned HTTP router responsible for handling all v1 API
// requests.
type Router struct {
	logger        zerolog.Logger
	db            *gorm.DB
	cookieCfg     gologin.CookieConfig
	sessionStore  *sessions.CookieStore
	oauth2Cfg     *oauth2.Config
	healthChecker *health.Health
	validate      *validator.Validate
}

func NewRouter(logger zerolog.Logger, db *gorm.DB, cookieCfg gologin.CookieConfig, sStore *sessions.CookieStore, oauth2Cfg *oauth2.Config) (*Router, error) {
	sqlDB, _ := db.DB()
	healthChecker, err := httputil.CreateHealthChecker(sqlDB, true)
	if err != nil {
		return nil, err
	}

	return &Router{
		db:            db,
		cookieCfg:     cookieCfg,
		sessionStore:  sStore,
		oauth2Cfg:     oauth2Cfg,
		healthChecker: healthChecker,
		validate:      validator.New(),
	}, nil
}

// Register registers all v1 HTTP handlers with the provided mux router and
// prefix path. All registered HTTP handlers come bundled with the appropriate
// middleware.
func (r *Router) Register(rtr *mux.Router, prefix string) {
	v1Router := rtr.PathPrefix(prefix).Subrouter()

	// build middleware chain
	mChain := middleware.Build(r.logger)

	// define and register the health endpoint
	v1Router.Handle(
		"/health",
		mChain.Then(handlers.NewJSONHandlerFunc(r.healthChecker, nil)),
	).Methods(httputil.MethodGET)

	// unauthenticated routes
	v1Router.Handle(
		"/modules/search",
		mChain.ThenFunc(r.SearchModules()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}", "q", "{q}").Methods(httputil.MethodGET)

	v1Router.Handle(
		"/modules",
		mChain.ThenFunc(r.GetAllModules()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(httputil.MethodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}",
		mChain.ThenFunc(r.GetModuleByID()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/versions",
		mChain.ThenFunc(r.GetModuleVersions()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/authors",
		mChain.ThenFunc(r.GetModuleAuthors()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/keywords",
		mChain.ThenFunc(r.GetModuleKeywords()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/users/{id:[0-9]+}",
		mChain.ThenFunc(r.GetUserByID()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/users/{id:[0-9]+}/modules",
		mChain.ThenFunc(r.GetUserModules()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/users",
		mChain.ThenFunc(r.GetAllUsers()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(httputil.MethodGET)

	v1Router.Handle(
		"/keywords",
		mChain.ThenFunc(r.GetAllKeywords()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(httputil.MethodGET)

	// authenticated routes
	v1Router.Handle(
		"/modules",
		mChain.ThenFunc(r.UpsertModule()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/user/tokens",
		mChain.ThenFunc(r.CreateUserToken()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/user/tokens",
		mChain.ThenFunc(r.GetUserTokens()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/user/tokens/{id:[0-9]+}",
		mChain.ThenFunc(r.RevokeUserToken()),
	).Methods(httputil.MethodDELETE)

	// session routes
	v1Router.Handle(
		"/session/start",
		mChain.Then(r.StartSession()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/session/authorize",
		mChain.Then(r.AuthorizeSession()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/session/logout",
		mChain.ThenFunc(r.LogoutSession()),
	).Methods(httputil.MethodPOST)
}

// UpsertModule implements a request handler to create or update a Cosmos SDK
// module.
// @Summary Publish a Cosmos SDK module
// @Tags modules
// @Accept  json
// @Produce  json
// @Param manifest body Manifest true "module manifest"
// @Success 200 {object} models.ModuleJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /modules [put]
func (r *Router) UpsertModule() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		var request Manifest
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("failed to read request: %w", err))
			return
		}

		if err := r.validate.Struct(request); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid request: %w", httputil.TransformValidationError(err)))
			return
		}

		module := ModuleFromManifest(request)

		// The publisher must already be an existing owner or must have accepted an
		// invitation by an existing owner.
		//
		// TODO: Handle invitations to allow other users to update existing modules.
		record, err := models.QueryModule(r.db, map[string]interface{}{"name": module.Name, "team": module.Team})
		if err == nil {
			// the module already exists so we check if the publisher is an owner
			var isOwner bool
			for i := 0; i < len(record.Owners) && !isOwner; i++ {
				if record.Owners[i].ID == authUser.ID {
					isOwner = true
				}
			}

			if !isOwner {
				httputil.RespondWithError(w, http.StatusBadRequest, errors.New("publisher must be an owner of the module"))
				return
			}

			module.Owners = record.Owners
		} else {
			module.Owners = []models.User{authUser}
		}

		module, err = module.Upsert(r.db)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, module)
	}
}

// GetModuleByID implements a request handler to retrieve a module by ID.
// @Summary Get a Cosmos SDK module by ID
// @Tags modules
// @Accept  json
// @Produce  json
// @Param id path int true "module ID"
// @Success 200 {object} models.ModuleJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /modules/{id} [get]
func (r *Router) GetModuleByID() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(r.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, module)
	}
}

// SearchModules implements a request handler to retrieve a set of module objects
// by search criteria.
// @Summary Search for Cosmos SDK modules by name, team, description and keywords
// @Tags modules
// @Accept  json
// @Produce  json
// @Param cursor query int true "pagination cursor"  default(0)
// @Param limit query int true "pagination limit"  default(100)
// @Param q query string true "search criteria"
// @Success 200 {object} httputil.PaginationResponse
// @Failure 400 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /modules/search [get]
func (r *Router) SearchModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := httputil.ParsePagination(req)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		query := req.URL.Query().Get("q")

		modules, err := models.SearchModules(r.db, query, cursor, limit)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := httputil.NewPaginationResponse(len(modules), limit, cursor, modules)
		httputil.RespondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetAllModules implements a request handler returning a paginated set of
// modules.
// @Summary Return a paginated set of all Cosmos SDK modules
// @Tags modules
// @Accept  json
// @Produce  json
// @Param cursor query int true "pagination cursor"  default(0)
// @Param limit query int true "pagination limit"  default(100)
// @Success 200 {object} httputil.PaginationResponse
// @Failure 400 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /modules [get]
func (r *Router) GetAllModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := httputil.ParsePagination(req)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		modules, err := models.GetAllModules(r.db, cursor, limit)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := httputil.NewPaginationResponse(len(modules), limit, cursor, modules)
		httputil.RespondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetModuleVersions implements a request handler to retrieve a module's set of
// versions by ID.
// @Summary Get all versions for a Cosmos SDK module by ID
// @Tags modules
// @Accept  json
// @Produce  json
// @Param id path int true "module ID"
// @Success 200 {array} models.ModuleVersionJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /modules/{id}/versions [get]
func (r *Router) GetModuleVersions() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(r.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, module.Versions)
	}
}

// GetModuleAuthors implements a request handler to retrieve a module's set of
// authors by ID.
// @Summary Get all authors for a Cosmos SDK module by ID
// @Tags modules
// @Accept  json
// @Produce  json
// @Param id path int true "module ID"
// @Success 200 {array} models.UserJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /modules/{id}/authors [get]
func (r *Router) GetModuleAuthors() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(r.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, module.Authors)
	}
}

// GetModuleKeywords implements a request handler to retrieve a module's set of
// keywords by ID.
// @Summary Get all keywords for a Cosmos SDK module by ID
// @Tags modules
// @Accept  json
// @Produce  json
// @Param id path int true "module ID"
// @Success 200 {array} models.KeywordJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /modules/{id}/keywords [get]
func (r *Router) GetModuleKeywords() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(r.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, module.Keywords)
	}
}

// GetUserByID implements a request handler to retrieve a user by ID.
// @Summary Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "user ID"
// @Success 200 {object} models.UserJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /users/{id} [get]
func (r *Router) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %w", err))
			return
		}

		user, err := models.GetUserByID(r.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, user)
	}
}

// GetAllUsers implements a request handler returning a paginated set of
// users.
// @Summary Return a paginated set of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Param cursor query int true "pagination cursor"  default(0)
// @Param limit query int true "pagination limit"  default(100)
// @Success 200 {object} httputil.PaginationResponse
// @Failure 400 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /users [get]
func (r *Router) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := httputil.ParsePagination(req)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		users, err := models.GetAllUsers(r.db, cursor, limit)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := httputil.NewPaginationResponse(len(users), limit, cursor, users)
		httputil.RespondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetUserModules implements a request handler to retrieve a set of modules
// authored by a given user by ID.
// @Summary Return a paginated set of all Cosmos SDK modules by user ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "user ID"
// @Success 200 {array} models.ModuleJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /users/{id}/modules [get]
func (r *Router) GetUserModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %w", err))
			return
		}

		modules, err := models.GetUserModules(r.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, modules)
	}
}

// CreateUserToken implements a request handler that creates a new API token for
// the authenticated user.
// @Summary Create a user API token
// @Tags users
// @Produce  json
// @Success 200 {object} models.UserTokenJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /user/tokens [put]
func (r *Router) CreateUserToken() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		numTokens := authUser.CountTokens(r.db)
		if numTokens >= MaxTokens {
			httputil.RespondWithError(w, http.StatusBadRequest, errors.New("maximum number of user API tokens reached"))
			return
		}

		token, err := authUser.CreateToken(r.db)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, token)
	}
}

// GetUserTokens implements a request handler returning all of an authenticated
// user's tokens.
// @Summary Get all API tokens by user ID
// @Tags users
// @Produce  json
// @Success 200 {array} models.UserTokenJSON
// @Failure 401 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /user/tokens [get]
func (r *Router) GetUserTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		tokens, err := authUser.GetTokens(r.db)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, tokens)
	}
}

// RevokeUserToken implements a request handler revoking a specific token from
// the authorized user.
// @Summary Revoke a user API token by ID
// @Tags users
// @Produce  json
// @Param id path int true "token ID"
// @Success 200 {object} models.UserTokenJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /user/tokens/{id} [delete]
func (r *Router) RevokeUserToken() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		token, err := models.QueryUserToken(r.db, map[string]interface{}{"id": id, "user_id": authUser.ID, "revoked": false})
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		token, err = token.Revoke(r.db)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, token)
	}
}

// GetAllKeywords implements a request handler returning a paginated set of
// keywords.
// @Summary Return a paginated set of all keywords
// @Tags keywords
// @Accept  json
// @Produce  json
// @Param cursor query int true "pagination cursor"  default(0)
// @Param limit query int true "pagination limit"  default(100)
// @Success 200 {object} httputil.PaginationResponse
// @Failure 400 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /keywords [get]
func (r *Router) GetAllKeywords() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := httputil.ParsePagination(req)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		keywords, err := models.GetAllKeywords(r.db, cursor, limit)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := httputil.NewPaginationResponse(len(keywords), limit, cursor, keywords)
		httputil.RespondWithJSON(w, http.StatusOK, paginated)
	}
}

// StartSession returns a request handler to begin a user session via Github
// OAuth authentication. The user must either grant or reject access. Upon
// granting access, Github will perform a callback where we create a session
// and obtain a token.
func (r *Router) StartSession() http.Handler {
	return github.StateHandler(r.cookieCfg, github.LoginHandler(r.oauth2Cfg, nil))
}

// AuthorizeSession returns a callback request handler for Github OAuth user
// authentication. After a user grants access, this callback handler will be
// executed. A session cookie will be saved and sent to the client. A user record
// will also be upserted.
func (r *Router) AuthorizeSession() http.Handler {
	return github.StateHandler(r.cookieCfg, github.CallbackHandler(r.oauth2Cfg, r.authorizeHandler(), nil))
}

func (r *Router) authorizeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		token, err := oauth2login.TokenFromContext(ctx)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get github token: %w", err))
			return
		}

		githubUser, err := github.UserFromContext(ctx)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get github user: %w", err))
			return
		}

		user := models.User{
			Name:              githubUser.GetLogin(),
			GithubUserID:      sql.NullInt64{Int64: githubUser.GetID(), Valid: true},
			GravatarID:        githubUser.GetGravatarID(),
			AvatarURL:         githubUser.GetAvatarURL(),
			GithubAccessToken: sql.NullString{String: token.AccessToken, Valid: true},
		}

		record, err := user.Upsert(r.db)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to upsert user: %w", err))
			return
		}

		session, err := r.sessionStore.Get(req, sessionName)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
			return
		}

		session.Values[sessionGithubID] = githubUser.GetID()
		session.Values[sessionUserID] = record.ID

		if err = session.Save(req, w); err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to save session: %w", err))
			return
		}

		http.Redirect(w, req, "/", http.StatusFound)
	}
}

// LogoutSession implements a request handler to terminate and logout of an
// existing session.
func (r *Router) LogoutSession() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := r.sessionStore.Get(req, sessionName)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
			return
		}

		// Remove session keys and set max age to -1 to trigger the deletion of the
		// cookie.
		delete(session.Values, sessionGithubID)
		delete(session.Values, sessionUserID)
		session.Options.MaxAge = -1

		if err = session.Save(req, w); err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to save session: %w", err))
			return
		}

		http.Redirect(w, req, "/", http.StatusFound)
	}
}

// authorize attempts to authorize the given request against the session cookie
// store or a bearer authorization header. If the session cookie does not exist,
// or the session has been deleted, or the supplied bearer authorization header
// is invalid, we treat the request as unauthorized and return false. Otherwise,
// we return the user record ID and true with no error to indicate successful
// authorization.
func (r *Router) authorize(req *http.Request) (models.User, bool, error) {
	session, err := r.sessionStore.Get(req, sessionName)
	if err != nil {
		return models.User{}, false, fmt.Errorf("failed to get session: %w", err)
	}

	var userID uint

	// check for a valid session cookie or bearer authorization header
	if v, ok := session.Values[sessionUserID]; ok {
		userID = v.(uint)
	} else if h := req.Header.Get("Authorization"); strings.HasPrefix(h, httputil.BearerSchema) {
		tokenStr := h[len(httputil.BearerSchema):]

		tokenUUID, err := uuid.FromString(tokenStr)
		if err != nil {
			return models.User{}, false, fmt.Errorf("failed to get parse token: %w", err)
		}

		token, err := models.QueryUserToken(r.db, map[string]interface{}{"token": tokenUUID.String(), "revoked": false})
		if err != nil {
			return models.User{}, false, err
		}

		token, err = token.IncrCount(r.db)
		if err != nil {
			return models.User{}, false, err
		}

		userID = token.UserID
	} else {
		return models.User{}, false, errors.New("unauthorized")
	}

	user, err := models.GetUserByID(r.db, userID)
	if err != nil {
		return models.User{}, false, err
	}

	return user, true, nil
}
