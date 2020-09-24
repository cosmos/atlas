package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/cosmos/atlas/server/models"
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
	db       *gorm.DB
	validate *validator.Validate
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		db:       db,
		validate: validator.New(),
	}
}

// GetModuleByID implements a request handler to retrieve a module by ID.
func (ctrl *Controller) GetModuleByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
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
	return func(w http.ResponseWriter, r *http.Request) {
		cursor, limit, err := parsePagination(r)
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
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
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
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
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
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
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
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
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
	return func(w http.ResponseWriter, r *http.Request) {
		cursor, limit, err := parsePagination(r)
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
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
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
	return func(w http.ResponseWriter, r *http.Request) {
		cursor, limit, err := parsePagination(r)
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
