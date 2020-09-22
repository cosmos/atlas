package server

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
