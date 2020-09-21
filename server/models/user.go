package models

import (
	"fmt"

	"gorm.io/gorm"
)

// User defines an entity that contributes to a Module type.
type User struct {
	gorm.Model

	Email             string `gorm:"not null;default:null" json:"email" yaml:"email"`
	Name              string `gorm:"not null;default:null" json:"name" yaml:"name"`
	URL               string `json:"url" yaml:"url"`
	AvatarURL         string `json:"avatar_url" yaml:"avatar_url"`
	GravatarID        string `json:"gravatar_id" yaml:"gravatar_id"`
	GithubAccessToken string `json:"-" yaml:"-"`
	APIToken          string `json:"-" yaml:"-"`

	Modules []Module `gorm:"many2many:module_authors" json:"modules" yaml:"modules"`
}

// GetUserByID returns a user by ID. If the user doesn't exist or if the
// query fails, an error is returned.
func GetUserByID(db *gorm.DB, id uint) (User, error) {
	var u User

	if err := db.First(&u, id).Error; err != nil {
		return User{}, fmt.Errorf("failed to query for user by ID: %w", err)
	}

	return u, nil
}
