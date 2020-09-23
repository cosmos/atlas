package server

import (
	"encoding/json"
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

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

// func (ctrl *Controller) UpsertModule(m Module) error {
// 	// err := ctrl.db.Model(&Module{}).Where("name = ?", m.Name).Updates(Module{})
// 	// if err != nil {
// 	// 	if gorm.Is(err) {
// 	// 		// db.Create(&newUser)  // create new record from newUser
// 	// 	}
// 	// }
// 	// if err := ctrl.validate.Struct(req); err != nil {
// 	// 	return nil, err
// 	// }
// }
