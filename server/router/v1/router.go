package v1

import (
	"database/sql"
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
	"github.com/microcosm-cc/bluemonday"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server/httputil"
	"github.com/cosmos/atlas/server/middleware"
	"github.com/cosmos/atlas/server/models"
)

const (
	sessionName        = "atlas_session"
	sessionGithubID    = "github_id"
	sessionUserID      = "user_Id"
	sessionRedirectURI = "redirect_uri"

	V1APIPathPrefix = "/api/v1"
)

var (
	// MaxTokens defines the maximum number of API tokens a user can create.
	MaxTokens int64 = 100

	paginationParams = []string{
		"page", "{page:[0-9]+}",
		"limit", "{limit:[0-9]+}",
	}
)

// Router implements a versioned HTTP router responsible for handling all v1 API
// requests
type Router struct {
	logger          zerolog.Logger
	cfg             config.Config
	db              *gorm.DB
	cookieCfg       gologin.CookieConfig
	sessionStore    *sessions.CookieStore
	oauth2Cfg       *oauth2.Config
	healthChecker   *health.Health
	validate        *validator.Validate
	sanitizer       Sanitizer
	ghClientCreator func(string) GitHubClientI
}

func NewRouter(
	logger zerolog.Logger,
	cfg config.Config,
	db *gorm.DB,
	cookieCfg gologin.CookieConfig,
	sStore *sessions.CookieStore,
	oauth2Cfg *oauth2.Config,
	ghClientCreator func(string) GitHubClientI,
) (*Router, error) {
	sqlDB, _ := db.DB()
	healthChecker, err := httputil.CreateHealthChecker(sqlDB, true)
	if err != nil {
		return nil, err
	}

	return &Router{
		logger:          logger,
		cfg:             cfg,
		db:              db,
		cookieCfg:       cookieCfg,
		sessionStore:    sStore,
		oauth2Cfg:       oauth2Cfg,
		healthChecker:   healthChecker,
		validate:        validator.New(),
		sanitizer:       newSanitizer(),
		ghClientCreator: ghClientCreator,
	}, nil
}

// Register registers all v1 HTTP handlers with the provided mux router and
// prefix path. All registered HTTP handlers come bundled with the appropriate
// middleware.
func (r *Router) Register(rtr *mux.Router, prefix string) {
	v1Router := rtr.PathPrefix(prefix).Subrouter()

	// handle all preflight request
	v1Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
	})

	// build middleware chain
	mChain := middleware.Build(r.logger, r.cfg)

	// define and register the health endpoint
	v1Router.Handle(
		"/health",
		mChain.Then(handlers.NewJSONHandlerFunc(r.healthChecker, nil)),
	).Methods(httputil.MethodGET)

	// ======================
	// unauthenticated routes
	// ======================

	v1Router.Handle(
		"/modules/search",
		mChain.ThenFunc(r.SearchModules()),
	).Queries(append(paginationParams, "q", "{q}")...).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/modules",
		mChain.ThenFunc(r.GetAllModules()),
	).Queries(paginationParams...).Methods(httputil.MethodGET)

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
		"/users/{name}",
		mChain.ThenFunc(r.GetUserByName()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/users/{name}/modules",
		mChain.ThenFunc(r.GetUserModules()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/users",
		mChain.ThenFunc(r.GetAllUsers()),
	).Queries(paginationParams...).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/keywords",
		mChain.ThenFunc(r.GetAllKeywords()),
	).Queries(paginationParams...).Methods(httputil.MethodGET)

	// ====================
	// authenticated routes
	// ====================

	v1Router.Handle(
		"/modules",
		mChain.ThenFunc(r.UpsertModule()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/star",
		mChain.ThenFunc(r.StarModule()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/unstar",
		mChain.ThenFunc(r.UnStarModule()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/me",
		mChain.ThenFunc(r.GetUser()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/me",
		mChain.ThenFunc(r.UpdateUser()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/me/confirm/{emailToken}",
		mChain.ThenFunc(r.ConfirmEmail()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/me/invite",
		mChain.ThenFunc(r.InviteOwner()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/me/invite/accept/{inviteToken}",
		mChain.ThenFunc(r.AcceptOwnerInvite()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/me/tokens",
		mChain.ThenFunc(r.CreateUserToken()),
	).Methods(httputil.MethodPUT)

	v1Router.Handle(
		"/me/tokens",
		mChain.ThenFunc(r.GetUserTokens()),
	).Methods(httputil.MethodGET)

	v1Router.Handle(
		"/me/tokens/{id:[0-9]+}",
		mChain.ThenFunc(r.RevokeUserToken()),
	).Methods(httputil.MethodDELETE)

	// ==============
	// session routes
	// ==============

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

// UpsertModule implements a request handler to publish a Cosmos SDK module.
// The authorized user is considered to be the publisher. The publisher must be
// an owner of the module and a contributor to the GitHub repository. If the
// module does not exist, the publisher is considered to be the first and only
// owner and subsequent owners may be invited by the publisher. An error is
// returned if the request body is invalid, the user is not authorized or if any
// database transaction fails.
//
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

		module := ModuleFromManifest(request, r.sanitizer)
		ghClient := r.ghClientCreator(authUser.GithubAccessToken.String)

		repo, err := ghClient.GetRepository(module.Version.Repo)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// set the module's team as the GitHub repository owner
		module.Team = repo.Owner

		// set the module's version publisher
		module.Version.PublishedBy = authUser.ID

		// verify the publisher is a contributor to the repository
		var isContributor bool
		for user := range repo.Contributors {
			if authUser.Name == user {
				isContributor = true
				break
			}
		}

		if !isContributor {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("publisher '%s' is not a contributor of this module", authUser.Name))
			return
		}

		// set the avatar URL for each author
		for i, author := range module.Authors {
			contributor, ok := repo.Contributors[author.Name]
			if ok {
				author.AvatarURL = contributor.GetAvatarURL()
				module.Authors[i] = author
			}
		}

		// The publisher must already be an existing owner or must have accepted an
		// invitation by an existing owner.
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
			// Otherwise, the module is new and we automatically assign the publisher
			// as the first and only owner.
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
//
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
//
// @Summary Search for Cosmos SDK modules by name, team, description and keywords
// @Tags modules
// @Accept  json
// @Produce  json
// @Param page query int true "pagination page"  default(1)
// @Param limit query int true "pagination limit"  default(100)
// @Param reverse query string false "pagination reverse"  default(false)
// @Param order query string false "pagination order by"  default(id)
// @Param q query string true "search criteria"
// @Success 200 {object} httputil.PaginationResponse
// @Failure 400 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /modules/search [get]
func (r *Router) SearchModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		pQuery, err := httputil.ParsePaginationQueryParams(req)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		query := req.URL.Query().Get("q")

		modules, paginator, err := models.SearchModules(r.db, query, pQuery)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := httputil.NewPaginationResponse(pQuery, paginator.PrevPage, paginator.NextPage, paginator.Total, modules)
		httputil.RespondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetAllModules implements a request handler returning a paginated set of
// modules.
//
// @Summary Return a paginated set of all Cosmos SDK modules
// @Tags modules
// @Accept  json
// @Produce  json
// @Param page query int true "pagination page"  default(1)
// @Param limit query int true "pagination limit"  default(100)
// @Param reverse query string false "pagination reverse"  default(false)
// @Param order query string false "pagination order by"  default(id)
// @Success 200 {object} httputil.PaginationResponse
// @Failure 400 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /modules [get]
func (r *Router) GetAllModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		pQuery, err := httputil.ParsePaginationQueryParams(req)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		modules, paginator, err := models.GetAllModules(r.db, pQuery)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := httputil.NewPaginationResponse(pQuery, paginator.PrevPage, paginator.NextPage, paginator.Total, modules)
		httputil.RespondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetModuleVersions implements a request handler to retrieve a module's set of
// versions by ID.
//
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
//
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
//
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

// GetUserByID implements a request handler to retrieve a user by name.
//
// @Summary Get a user by name
// @Tags users
// @Accept  json
// @Produce  json
// @Param name path string true "user name"
// @Success 200 {object} models.UserJSON
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /users/{name} [get]
func (r *Router) GetUserByName() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)

		user, err := models.QueryUser(r.db, map[string]interface{}{"name": params["name"]})
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
//
// @Summary Return a paginated set of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Param page query int true "pagination page"  default(1)
// @Param limit query int true "pagination limit"  default(100)
// @Param reverse query string false "pagination reverse"  default(false)
// @Param order query string false "pagination order by"  default(id)
// @Success 200 {object} httputil.PaginationResponse
// @Failure 400 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /users [get]
func (r *Router) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		pQuery, err := httputil.ParsePaginationQueryParams(req)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		users, paginator, err := models.GetAllUsers(r.db, pQuery)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := httputil.NewPaginationResponse(pQuery, paginator.PrevPage, paginator.NextPage, paginator.Total, users)
		httputil.RespondWithJSON(w, http.StatusOK, paginated)
	}
}

// GetUserModules implements a request handler to retrieve a set of modules
// authored by a given user by name.
//
// @Summary Return a set of all Cosmos SDK modules published by a given user
// @Tags users
// @Accept  json
// @Produce  json
// @Param name path string true "user name"
// @Success 200 {array} models.ModuleJSON
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /users/{name}/modules [get]
func (r *Router) GetUserModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)

		modules, err := models.GetUserModules(r.db, params["name"])
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

// InviteOwner implements a request handler to invite a user to be an owner of a
// module.
//
// @Summary Invite a user to be an owner of a module
// @Tags users
// @Produce  json
// @Accept  json
// @Param invite body ModuleInvite true "invitation"
// @Success 200 {object} boolean
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /me/invite [put]
func (r *Router) InviteOwner() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		var requestBody ModuleInvite
		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("failed to read request: %w", err))
			return
		}

		if err := r.validate.Struct(requestBody); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid request: %w", httputil.TransformValidationError(err)))
			return
		}

		module, err := models.GetModuleByID(r.db, requestBody.ModuleID)
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		// ensure invitee is not already an owner
		for _, o := range module.Owners {
			if o.Name == requestBody.User {
				httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("'%s' is already an owner", requestBody.User))
				return
			}
		}

		invitee, err := models.QueryUser(r.db, map[string]interface{}{"name": requestBody.User})
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		// ensure invitee has a verified email
		if !invitee.EmailConfirmed {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("'%s' must confirm their email address", requestBody.User))
			return
		}

		// upsert invite record
		moi, err := models.ModuleOwnerInvite{ModuleID: module.ID, InvitedByUserID: authUser.ID, InvitedUserID: invitee.ID}.Upsert(r.db)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		// send invite
		acceptURL := fmt.Sprintf("%s/accept/%s", r.cfg.String(config.DomainName), moi.Token)
		if err := r.sendOwnerInvitation(acceptURL, authUser.Name, invitee, module); err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, true)
	}
}

// AcceptOwnerInvite implements a request handler for accepting a module owner
// invitation.
//
// @Summary Accept a module owner invitation
// @Tags users
// @Produce  json
// @Param inviteToken path string true "invite token"
// @Success 200 {object} models.ModuleJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /me/invite/accept/{inviteToken} [put]
func (r *Router) AcceptOwnerInvite() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		inviteToken := mux.Vars(req)["inviteToken"]
		moi, err := models.QueryModuleOwnerInvite(r.db, map[string]interface{}{"invited_user_id": authUser.ID, "token": inviteToken})
		if err != nil {
			httputil.RespondWithError(w, http.StatusNotFound, err)
			return
		}

		// prevent stale invites from being accepted
		if time.Since(moi.UpdatedAt) > 24*time.Hour {
			httputil.RespondWithError(w, http.StatusBadRequest, errors.New("expired module owner invitation"))
			return
		}

		module, err := models.GetModuleByID(r.db, moi.ModuleID)
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		module, err = module.AddOwner(r.db, authUser)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, module)
	}
}

// CreateUserToken implements a request handler that creates a new API token for
// the authenticated user.
//
// @Summary Create a user API token
// @Tags users
// @Produce  json
// @Param token body Token true "token name"
// @Success 200 {object} models.UserTokenJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /me/tokens [put]
func (r *Router) CreateUserToken() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		var request Token
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("failed to read request: %w", err))
			return
		}

		if err := r.validate.Struct(request); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid request: %w", httputil.TransformValidationError(err)))
			return
		}

		numTokens := authUser.CountTokens(r.db)
		if numTokens >= MaxTokens {
			httputil.RespondWithError(w, http.StatusBadRequest, errors.New("maximum number of user API tokens reached"))
			return
		}

		token, err := authUser.CreateToken(r.db, request.Name)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, token)
	}
}

// GetUserTokens implements a request handler returning all of an authenticated
// user's tokens.
//
// @Summary Get all API tokens by user ID
// @Tags users
// @Produce  json
// @Success 200 {array} models.UserTokenJSON
// @Failure 401 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /me/tokens [get]
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
//
// @Summary Revoke a user API token by ID
// @Tags users
// @Produce  json
// @Param id path int true "token ID"
// @Success 200 {object} models.UserTokenJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /me/tokens/{id} [delete]
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
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %w", err))
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

// StarModule implements a request handler for adding a favorite by a user to a
// given module.
//
// @Summary Add a favorite for a module
// @Tags modules
// @Produce  json
// @Param id path int true "module ID"
// @Success 200 {object} ModuleStars
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /modules/{id}/star [put]
func (r *Router) StarModule() http.HandlerFunc {
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

		module, err := models.GetModuleByID(r.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		stars, err := module.Star(r.db, authUser.ID)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, ModuleStars{Stars: stars})
	}
}

// UnStarModule implements a request handler for removing a favorite by a user
// to a given module.
//
// @Summary Remove a favorite for a module
// @Tags modules
// @Produce  json
// @Param id path int true "module ID"
// @Success 200 {object} ModuleStars
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /modules/{id}/unstar [put]
func (r *Router) UnStarModule() http.HandlerFunc {
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

		module, err := models.GetModuleByID(r.db, uint(id))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, gorm.ErrRecordNotFound) {
				code = http.StatusNotFound
			}

			httputil.RespondWithError(w, code, err)
			return
		}

		stars, err := module.UnStar(r.db, authUser.ID)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, ModuleStars{Stars: stars})
	}
}

// GetUser returns the current authenticated user.
//
// @Summary Get the current authenticated user
// @Tags users
// @Produce  json
// @Success 200 {object} models.UserJSON
// @Failure 401 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /me [get]
func (r *Router) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, authUser)
	}
}

// UpdateUser updates an existing user record.
//
// @Summary Update the current authenticated user
// @Tags users
// @Produce  json
// @Param user body User true "user"
// @Success 200 {object} boolean
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Security APIKeyAuth
// @Router /me [put]
func (r *Router) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		var request User
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("failed to read request: %w", err))
			return
		}

		if err := r.validate.Struct(request); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid request: %w", httputil.TransformValidationError(err)))
			return
		}

		if request.Email == authUser.Email.String && authUser.EmailConfirmed {
			httputil.RespondWithError(w, http.StatusBadRequest, errors.New("email already confirmed"))
			return
		}

		if request.Email != authUser.Email.String {
			authUser.EmailConfirmed = false
		}

		// If the email is non-empty and requires confirmation, either because it is
		// new or it has been updated, we send an email confirmation.
		if !authUser.EmailConfirmed && request.Email != "" {
			uec, err := models.UserEmailConfirmation{UserID: authUser.ID, Email: request.Email}.Upsert(r.db)
			if err != nil {
				httputil.RespondWithError(w, http.StatusInternalServerError, err)
				return
			}

			confirmURL := fmt.Sprintf("%s/confirm/%s", r.cfg.String(config.DomainName), uec.Token)

			if err := r.sendEmailConfirmation(authUser.Name, request.Email, confirmURL); err != nil {
				httputil.RespondWithError(w, http.StatusInternalServerError, err)
				return
			}
		}

		if _, err := authUser.Upsert(r.db); err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, true)
	}
}

// ConfirmEmail implements a request handler for confirming a user email address.
//
// @Summary Confirm a user email confirmation
// @Tags users
// @Produce  json
// @Param emailToken path string true "email token"
// @Success 200 {object} models.UserJSON
// @Failure 400 {object} httputil.ErrResponse
// @Failure 401 {object} httputil.ErrResponse
// @Failure 404 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /me/confirm/{emailToken} [put]
func (r *Router) ConfirmEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authUser, ok, err := r.authorize(req)
		if err != nil || !ok {
			httputil.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		emailToken := mux.Vars(req)["emailToken"]
		uec, err := models.QueryUserEmailConfirmation(r.db, map[string]interface{}{"user_id": authUser.ID, "token": emailToken})
		if err != nil {
			httputil.RespondWithError(w, http.StatusNotFound, err)
			return
		}

		// prevent stale confirmations from being accepted
		if time.Since(uec.UpdatedAt) > 10*time.Minute {
			httputil.RespondWithError(w, http.StatusBadRequest, errors.New("expired email confirmation"))
			return
		}

		user, err := authUser.ConfirmEmail(r.db, uec)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		httputil.RespondWithJSON(w, http.StatusOK, user)
	}
}

// GetAllKeywords implements a request handler returning a paginated set of
// keywords.
//
// @Summary Return a paginated set of all keywords
// @Tags keywords
// @Accept  json
// @Produce  json
// @Param page query int true "pagination page"  default(1)
// @Param limit query int true "pagination limit"  default(100)
// @Param reverse query string false "pagination reverse"  default(false)
// @Param order query string false "pagination order by"  default(id)
// @Success 200 {object} httputil.PaginationResponse
// @Failure 400 {object} httputil.ErrResponse
// @Failure 500 {object} httputil.ErrResponse
// @Router /keywords [get]
func (r *Router) GetAllKeywords() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		pQuery, err := httputil.ParsePaginationQueryParams(req)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		keywords, paginator, err := models.GetAllKeywords(r.db, pQuery)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		paginated := httputil.NewPaginationResponse(pQuery, paginator.PrevPage, paginator.NextPage, paginator.Total, keywords)
		httputil.RespondWithJSON(w, http.StatusOK, paginated)
	}
}

// StartSession returns a request handler to begin a user session via Github
// OAuth authentication. The user must either grant or reject access. Upon
// granting access, Github will perform a callback where we create a session
// and obtain a token.
func (r *Router) StartSession() http.Handler {
	loginHandler := func(w http.ResponseWriter, req *http.Request) {
		req.Header.Set("Access-Control-Allow-Origin", "*")

		ctx := req.Context()

		state, err := oauth2login.StateFromContext(ctx)
		if err != nil {
			ctx = gologin.WithError(ctx, err)
			gologin.DefaultFailureHandler.ServeHTTP(w, req.WithContext(ctx))
			return
		}

		// Create and store a session with the client's referrer address so we know
		// where to redirect to after authentication is complete.
		session, err := r.sessionStore.Get(req, sessionName)
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
			return
		}

		session.Values[sessionRedirectURI] = req.Referer()

		if err = session.Save(req, w); err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to save session: %w", err))
			return
		}

		authURL := r.oauth2Cfg.AuthCodeURL(state)
		http.Redirect(w, req, authURL, http.StatusFound)
	}

	return github.StateHandler(r.cookieCfg, http.HandlerFunc(loginHandler))
}

// AuthorizeSession returns a callback request handler for Github OAuth user
// authentication. After a user grants access, this callback handler will be
// executed. A session cookie will be saved and sent to the client. A user record
// will also be upserted.
func (r *Router) AuthorizeSession() http.Handler {
	return github.StateHandler(
		r.cookieCfg,
		github.CallbackHandler(
			r.oauth2Cfg,
			r.authorizeHandler(),
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				session, err := r.sessionStore.Get(req, sessionName)
				if err != nil {
					httputil.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
					return
				}

				referrer, ok := session.Values[sessionRedirectURI].(string)
				if !ok || referrer == "" {
					referrer = "/"
				}

				http.Redirect(w, req, referrer, http.StatusFound)
			}),
		),
	)
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
			FullName:          githubUser.GetName(),
			GithubUserID:      models.NewNullInt64(githubUser.GetID()),
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

		referrer, ok := session.Values[sessionRedirectURI].(string)
		if !ok || referrer == "" {
			referrer = "/"
		}

		http.Redirect(w, req, referrer, http.StatusFound)
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

func newSanitizer() Sanitizer {
	return bluemonday.NewPolicy().
		RequireParseableURLs(true).
		AllowRelativeURLs(false).
		AllowURLSchemes("http", "https")
}
