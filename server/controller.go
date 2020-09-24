package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dghubble/gologin"
	"github.com/dghubble/gologin/github"
	oauth2Login "github.com/dghubble/gologin/oauth2"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	githuboauth2 "golang.org/x/oauth2/github"
	"gorm.io/gorm"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server/models"
)

const (
	sessionName      = "atlas_session"
	sessionUserKey   = "github_id"
	sessionUserLogin = "github_login"
)

// type (
// 	moduleRequest struct {
// 		Name        string        `json:"name" yaml:"name" validate:"required"`
// 		Description string        `json:"description" yaml:"description"`
// 		Homepage    string        `json:"homepage" yaml:"homepage" validate:"url"`
// 		Repo        string        `json:"repo" yaml:"repo" validate:"url"`
// 		Version     ModuleVersion `json:"version" yaml:"version" validate:"required"`
// 		BugTracker  BugTracker    `json:"bug_tracker" yaml:"bug_tracker"`
// 		Keywords    []Keyword     `json:"keywords" yaml:"keywords"`
// 	}
// )

// Controller contains a wrapper around a Database and is responsible for
// implementing API request handlers.
type Controller struct {
	cookieCfg    gologin.CookieConfig
	sessionStore *sessions.CookieStore
	db           *gorm.DB
	oauth2Cfg    *oauth2.Config
	validate     *validator.Validate
}

func NewController(db *gorm.DB, cfg config.Config) *Controller {
	cookieCfg := gologin.DefaultCookieConfig
	sessionStore := sessions.NewCookieStore([]byte(cfg.String(config.FlagSessionKey)), nil)
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = true
	sessionStore.Options.MaxAge = 3600 * 24 * 7 // 1 week

	if cfg.Bool(config.FlagDev) {
		cookieCfg = gologin.DebugOnlyCookieConfig
		sessionStore.Options.Secure = false
	}

	return &Controller{
		cookieCfg:    cookieCfg,
		sessionStore: sessionStore,
		db:           db,
		validate:     validator.New(),
		oauth2Cfg: &oauth2.Config{
			ClientID:     cfg.String(config.FlagGHClientID),
			ClientSecret: cfg.String(config.FlagGHClientSecret),
			RedirectURL:  cfg.String(config.FlagGHRedirectURL),
			Endpoint:     githuboauth2.Endpoint,
		},
	}
}

// GetModuleByID implements a request handler to retrieve a module by ID.
func (ctrl *Controller) GetModuleByID() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(ctrl.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module)
	}
}

// GetAllModules implements a request handler returning a paginated set of
// modules.
func (ctrl *Controller) GetAllModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := parsePagination(req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		modules, err := models.GetAllModules(ctrl.db, cursor, limit)
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
func (ctrl *Controller) GetModuleVersions() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(ctrl.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module.Versions)
	}
}

// GetModuleAuthors implements a request handler to retreive a module's set of
// authors by ID.
func (ctrl *Controller) GetModuleAuthors() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(ctrl.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module.Authors)
	}
}

// GetModuleKeywords implements a request handler to retreive a module's set of
// keywords by ID.
func (ctrl *Controller) GetModuleKeywords() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid module ID: %w", err))
			return
		}

		module, err := models.GetModuleByID(ctrl.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, module.Keywords)
	}
}

// GetUserByID implements a request handler to retrieve a user by ID.
func (ctrl *Controller) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %w", err))
			return
		}

		user, err := models.GetUserByID(ctrl.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, user)
	}
}

// GetAllUsers implements a request handler returning a paginated set of
// users.
func (ctrl *Controller) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := parsePagination(req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		users, err := models.GetAllUsers(ctrl.db, cursor, limit)
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
func (ctrl *Controller) GetUserModules() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		idStr := params["id"]

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %w", err))
			return
		}

		modules, err := models.GetUserModules(ctrl.db, uint(id))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, modules)
	}
}

// GetAllKeywords implements a request handler returning a paginated set of
// keywords.
func (ctrl *Controller) GetAllKeywords() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cursor, limit, err := parsePagination(req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		keywords, err := models.GetAllKeywords(ctrl.db, cursor, limit)
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
func (ctrl *Controller) BeginSession() http.Handler {
	return github.StateHandler(ctrl.cookieCfg, github.LoginHandler(ctrl.oauth2Cfg, nil))
}

// AuthorizeSession returns a callback request handler for Github OAuth user
// authentication. After a user grants access, this callback handler will be
// executed. A session cookie will be saved and sent to the client. A user record
// will also be upserted.
func (ctrl *Controller) AuthorizeSession() http.Handler {
	return github.StateHandler(ctrl.cookieCfg, github.CallbackHandler(ctrl.oauth2Cfg, ctrl.authorizeSession(), nil))
}

func (ctrl *Controller) authorizeSession() http.Handler {
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
			GithubAccessToken: token.AccessToken,
			GravatarID:        githubUser.GetGravatarID(),
			AvatarURL:         githubUser.GetAvatarURL(),
		}
		if _, err := user.Upsert(ctrl.db); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to upsert user: %w", err))
			return
		}

		session, err := ctrl.sessionStore.Get(req, sessionName)
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

	return http.HandlerFunc(fn)
}

// LogoutSession implements a request handler to terminate and logout of an
// existing session.
func (ctrl *Controller) LogoutSession() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := ctrl.sessionStore.Get(req, sessionName)
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
