package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User defines an entity that contributes to a Module type.
type User struct {
	GormModel

	Email             string `json:"email" yaml:"email"`
	Name              string `json:"name" yaml:"name"`
	URL               string `json:"url" yaml:"url"`
	AvatarURL         string `json:"avatar_url" yaml:"avatar_url"`
	GravatarID        string `json:"gravatar_id" yaml:"gravatar_id"`
	GithubAccessToken string `json:"-" yaml:"-"`
	APIToken          string `json:"-" yaml:"-"`

	Modules []Module `gorm:"many2many:module_authors" json:"-" yaml:"-"`
}

// GetUserByID returns a User by ID. If the user doesn't exist or if the
// query fails, an error is returned.
func GetUserByID(db *gorm.DB, id uint) (User, error) {
	var u User

	if err := db.Preload(clause.Associations).First(&u, id).Error; err != nil {
		return User{}, fmt.Errorf("failed to query for user by ID: %w", err)
	}

	return u, nil
}

// GetUserModules returns a set of Module's authored by a given User by ID.
func GetUserModules(db *gorm.DB, id uint) ([]Module, error) {
	user, err := GetUserByID(db, id)
	if err != nil {
		return []Module{}, err
	}

	moduleIDs := make([]uint, len(user.Modules))
	for i, mod := range user.Modules {
		moduleIDs[i] = mod.ID
	}

	var modules []Module
	if err := db.Preload(clause.Associations).Where(moduleIDs).Find(&modules).Error; err != nil {
		return []Module{}, fmt.Errorf("failed to query for user by ID: %w", err)
	}

	return modules, nil
}

// Query performs a query for a User record where the search criteria is defined
// by the receiver object. The resulting record, if it exists, is returned. If
// the query fails or the record does not exist, an error is returned.
func (u User) Query(db *gorm.DB) (User, error) {
	var record User

	if err := db.Where(u).First(&record).Error; err != nil {
		return User{}, fmt.Errorf("failed to query user: %w", err)
	}

	return record, nil
}

// GetAllUsers returns a slice of User objects paginated by a cursor and a
// limit. The cursor must be the ID of the last retrieved object. An error is
// returned upon database query failure.
func GetAllUsers(db *gorm.DB, cursor uint, limit int) ([]User, error) {
	var users []User

	if err := db.Limit(limit).Order("id asc").Where("id > ?", cursor).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to query for users: %w", err)
	}

	return users, nil
}
