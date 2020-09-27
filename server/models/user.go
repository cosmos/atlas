package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User defines an entity that contributes to a Module type.
type User struct {
	gorm.Model

	Name              string         `json:"name" yaml:"name"`
	GithubUserID      sql.NullInt64  `json:"-" yaml:"-"`
	GithubAccessToken sql.NullString `json:"-" yaml:"-"`
	Email             sql.NullString `json:"-" yaml:"-"`
	URL               string         `json:"url" yaml:"url"`
	AvatarURL         string         `json:"avatar_url" yaml:"avatar_url"`
	GravatarID        string         `json:"gravatar_id" yaml:"gravatar_id"`

	Modules []Module `gorm:"many2many:module_authors" json:"-" yaml:"-"`
}

// MarshalJSON implements custom JSON marshaling for the User model.
func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		GormModelJSON

		Name       string `json:"name" yaml:"name"`
		URL        string `json:"url" yaml:"url"`
		AvatarURL  string `json:"avatar_url" yaml:"avatar_url"`
		GravatarID string `json:"gravatar_id" yaml:"gravatar_id"`
	}{
		GormModelJSON: GormModelJSON{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Name:       u.Name,
		URL:        u.URL,
		AvatarURL:  u.AvatarURL,
		GravatarID: u.GravatarID,
	})
}

// Upsert creates or updates a User record. Note, this should only be called
// when authenticating a user. When authors are associated with a Module, they
// are either fetched or created by their name and email.
func (u User) Upsert(db *gorm.DB) (User, error) {
	var record User

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("name = ?", u.Name).First(&record).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := tx.Create(&u).Error; err != nil {
					return fmt.Errorf("failed to create user: %w", err)
				}

				// commit the tx
				return nil
			} else {
				return fmt.Errorf("failed to query for user: %w", err)
			}
		}

		if err := tx.Model(&record).Updates(User{
			Name:              u.Name,
			GithubUserID:      u.GithubUserID,
			GithubAccessToken: u.GithubAccessToken,
			AvatarURL:         u.AvatarURL,
			GravatarID:        u.GravatarID,
		}).Error; err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		// commit the tx
		return nil
	})
	if err != nil {
		return User{}, err
	}

	return GetUserByID(db, record.ID)
}

// GetUserByID returns a User by ID. If the user doesn't exist or if the
// query fails, an error is returned.
func GetUserByID(db *gorm.DB, id uint) (User, error) {
	var u User

	if err := db.First(&u, id).Error; err != nil {
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
