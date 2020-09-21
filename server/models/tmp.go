package models

import (
	"gorm.io/gorm"
)

// Keyword defines a module keyword, where a module can have one or more keywords.
type Keyword struct {
	gorm.Model

	Name string `json:"name" yaml:"name"`
}

// ModuleVersion defines a version associated with a unique module.
type ModuleVersion struct {
	gorm.Model

	Version  string `json:"version" yaml:"version"`
	ModuleID uint   `json:"-" yaml:"-"`
}

// User defines an entity that contributes to a Module type.
type User struct {
	gorm.Model

	Name              string `json:"name" yaml:"name"`
	URL               string `json:"url" yaml:"url"`
	Email             string `json:"email" yaml:"email"`
	GithubAccessToken string `json:"github_access_token" yaml:"github_access_token"`
	APIToken          string `json:"api_token" yaml:"api_token"`
	AvatarURL         string `json:"avatar_url" yaml:"avatar_url"`
	GravatarID        string `json:"gravatar_id" yaml:"gravatar_id"`

	Modules []Module `gorm:"many2many:module_authors" json:"modules" yaml:"modules"`
}

// ModuleKeywords defines the type relationship between a module and all the
// associated keywords.
type ModuleKeywords struct {
	ModuleID  uint
	KeywordID uint
}

// ModuleAuthors defines the type relationship between a module and all the
// associated authors.
type ModuleAuthors struct {
	ModuleID uint
	UserID   uint
}

// BugTracker defines the metadata information for reporting bug reports on a
// given Module type.
type BugTracker struct {
	gorm.Model

	URL      string `gorm:"not null;default:null" json:"url" yaml:"url"`
	Contact  string `gorm:"not null;default:null" json:"contact" yaml:"contact"`
	ModuleID uint
}
