package server

import "github.com/cosmos/atlas/server/models"

// PaginationResponse defines a generic type encapsulating a paginated response.
// Client should not rely on decoding into this type as the Results is an
// interface.
type PaginationResponse struct {
	Limit   int         `json:"limit" yaml:"limit"`
	Cursor  uint        `json:"cursor" yaml:"cursor"`
	Count   int         `json:"count" yaml:"count"`
	Results interface{} `json:"results" yaml:"results"`
}

func NewPaginationResponse(count, limit int, cursor uint, results interface{}) PaginationResponse {
	return PaginationResponse{
		Limit:   limit,
		Cursor:  cursor,
		Count:   count,
		Results: results,
	}
}

// UserRequest defines a type wrapper for defining User model data in a request.
type UserRequest struct {
	Name  string `json:"name" yaml:"name" validate:"required"`
	Email string `json:"email" yaml:"email" validate:"omitempty,email"`
}

// KeywordRequest defines a type wrapper for defining Keyword model data in a
// request.
type KeywordRequest struct {
	Name string `json:"name" yaml:"name" validate:"omitempty,required"`
}

// BugTrackerRequest defines a type wrapper for defining BugTracker model data
// in a request.
type BugTrackerRequest struct {
	URL     string `json:"url" yaml:"url" validate:"required,url"`
	Contact string `json:"contact" yaml:"contact" validate:"required,email"`
}

// ModuleRequest defines a type wrapper for defining Module model data in a
// request.
type ModuleRequest struct {
	Name        string             `json:"name" yaml:"name" validate:"required"`
	Team        string             `json:"team" yaml:"team" validate:"required"`
	Repo        string             `json:"repo" yaml:"repo" validate:"required,url"`
	Version     string             `json:"version" yaml:"version" validate:"required"`
	Authors     []UserRequest      `json:"authors" yaml:"authors" validate:"required,dive"`
	Description string             `json:"description" yaml:"description"`
	Homepage    string             `json:"homepage" yaml:"homepage" validate:"omitempty,url"`
	BugTracker  *BugTrackerRequest `json:"bug_tracker" yaml:"bug_tracker" validate:"omitempty,dive"`
	Keywords    []KeywordRequest   `json:"keywords" yaml:"keywords" validate:"dive"`
}

// ModuleFromRequest converts a ModuleRequest to a Module model.
func ModuleFromRequest(req ModuleRequest) models.Module {
	authors := make([]models.User, len(req.Authors))
	for i, a := range authors {
		authors[i] = models.User{Name: a.Name, Email: a.Email}
	}

	keywords := make([]models.Keyword, len(req.Keywords))
	for i, k := range req.Keywords {
		keywords[i] = models.Keyword{Name: k.Name}
	}

	bugTracker := models.BugTracker{}
	if req.BugTracker != nil {
		bugTracker = models.BugTracker{URL: req.BugTracker.URL, Contact: req.BugTracker.Contact}
	}

	return models.Module{
		Name:        req.Name,
		Team:        req.Team,
		Repo:        req.Repo,
		Version:     req.Version,
		Description: req.Description,
		Homepage:    req.Homepage,
		Authors:     authors,
		Keywords:    keywords,
		BugTracker:  bugTracker,
	}
}
