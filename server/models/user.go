package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	// UserToken defines a user created API token.
	UserToken struct {
		gorm.Model

		UserID  uint      `json:"user_id" yaml:"user_id"`
		Token   uuid.UUID `json:"token" yaml:"token"`
		Revoked bool      `json:"revoked" yaml:"revoked"`
		Count   uint      `json:"count" yaml:"count"`
	}

	// User defines an entity that contributes to a Module type.
	User struct {
		gorm.Model

		Name              string         `json:"name" yaml:"name"`
		GithubUserID      sql.NullInt64  `json:"-" yaml:"-"`
		GithubAccessToken sql.NullString `json:"-" yaml:"-"`
		Email             sql.NullString `json:"-" yaml:"-"`
		URL               string         `json:"url" yaml:"url"`
		AvatarURL         string         `json:"avatar_url" yaml:"avatar_url"`
		GravatarID        string         `json:"gravatar_id" yaml:"gravatar_id"`

		// many-to-many relationships
		Modules []Module `gorm:"many2many:module_authors" json:"-" yaml:"-"`

		// one-to-many relationships
		Tokens []UserToken `gorm:"foreignKey:user_id" json:"-" yaml:"-"`
	}
)

// MarshalJSON implements custom JSON marshaling for the UserToken model.
func (ut UserToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		GormModelJSON

		UserID  uint      `json:"user_id" yaml:"user_id"`
		Token   uuid.UUID `json:"token" yaml:"token"`
		Revoked bool      `json:"revoked" yaml:"revoked"`
		Count   uint      `json:"count" yaml:"count"`
	}{
		GormModelJSON: GormModelJSON{
			ID:        ut.ID,
			CreatedAt: ut.CreatedAt,
			UpdatedAt: ut.UpdatedAt,
		},
		UserID:  ut.UserID,
		Token:   ut.Token,
		Revoked: ut.Revoked,
		Count:   ut.Count,
	})
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
			Email:             u.Email,
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

	return QueryUser(db, map[string]interface{}{"name": u.Name})
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

// QueryUser performs a query for a User record. The resulting record, if it exists,
// is returned. If the query fails or the record does not exist, an error is
// returned.
func QueryUser(db *gorm.DB, query map[string]interface{}) (User, error) {
	var record User

	if err := db.Where(query).First(&record).Error; err != nil {
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

// Revoke revokes a token. It returns an error upon failure.
func (ut UserToken) Revoke(db *gorm.DB) (UserToken, error) {
	if err := db.Model(&ut).Updates(UserToken{
		Revoked: true,
	}).Error; err != nil {
		return UserToken{}, fmt.Errorf("failed to revoke user token: %w", err)
	}

	return ut, nil
}

// IncrCount increments a token's count. It returns an error upon failure.
func (ut UserToken) IncrCount(db *gorm.DB) (UserToken, error) {
	if err := db.Model(&ut).Updates(UserToken{
		Count: ut.Count + 1,
	}).Error; err != nil {
		return UserToken{}, fmt.Errorf("failed to increment user token count: %w", err)
	}

	return ut, nil
}

// BeforeCreate will create and set the UserToken UUID.
func (ut *UserToken) BeforeCreate(tx *gorm.DB) error {
	ut.Token = uuid.NewV4()
	return nil
}

// QueryUserToken performs a query for a UserToken record. The resulting record,
// if it exists, is returned. If the query fails or the record does not exist,
// an error is returned.
func QueryUserToken(db *gorm.DB, query map[string]interface{}) (UserToken, error) {
	var record UserToken

	if err := db.Where(query).First(&record).Error; err != nil {
		return UserToken{}, fmt.Errorf("failed to query user token: %w", err)
	}

	return record, nil
}

// CreateToken creates a new UserToken for a given User model. It returns an
// error upon failure.
func (u User) CreateToken(db *gorm.DB) (UserToken, error) {
	token := UserToken{UserID: u.ID}

	// Note: The Append call will create a new UserToken record.
	if err := db.Model(&u).Association("Tokens").Append(&token); err != nil {
		return UserToken{}, fmt.Errorf("failed to assign token to user: %w", err)
	}

	return token, nil
}

// GetTokens returns all UserToken records for a given User record. It returns
// an error upon failure.
func (u User) GetTokens(db *gorm.DB) ([]UserToken, error) {
	var tokens []UserToken

	if err := db.Model(&u).Association("Tokens").Find(&tokens); err != nil {
		return nil, fmt.Errorf("failed to fetch user tokens: %w", err)
	}

	return tokens, nil
}

// CountTokens returns the total number of API tokens belonging to a User.
func (u User) CountTokens(db *gorm.DB) int64 {
	return db.Model(&u).Association("Tokens").Count()
}
