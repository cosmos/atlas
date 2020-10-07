package server

import (
	"github.com/cosmos/atlas/server/models"
)

// PaginationResponse defines a generic type encapsulating a paginated response.
// Client should not rely on decoding into this type as the Results is an
// interface.
type PaginationResponse struct {
	Limit   int         `json:"limit"`
	Cursor  uint        `json:"cursor"`
	Count   int         `json:"count"`
	Results interface{} `json:"results"`
}

func NewPaginationResponse(count, limit int, cursor uint, results interface{}) PaginationResponse {
	return PaginationResponse{
		Limit:   limit,
		Cursor:  cursor,
		Count:   count,
		Results: results,
	}
}

type (
	// ModuleManifest defines the primary module fields in a module's manifest.
	ModuleManifest struct {
		Name          string   `json:"name" toml:"name" validate:"required"`
		Team          string   `json:"team" toml:"team" validate:"required"`
		Repo          string   `json:"repo" toml:"repo" validate:"required,url"`
		Keywords      []string `json:"keywords" toml:"keywords" validate:"omitempty,gt=0,unique,dive,gt=0"`
		Description   string   `json:"description" toml:"description"`
		Homepage      string   `json:"homepage" toml:"homepage" validate:"omitempty,url"`
		Documentation string   `json:"documentation" toml:"documentation" validate:"omitempty,url"`
	}

	// AuthorsManifest defines author information in a module's manifest.
	AuthorsManifest struct {
		Name  string `json:"name" toml:"name" validate:"required"`
		Email string `json:"email" toml:"email" validate:"omitempty,email"`
	}

	// BugTackerManifest defines the bug tracker information in a module's manifest.
	BugTackerManifest struct {
		URL     string `json:"url" toml:"url" validate:"omitempty,url"`
		Contact string `json:"contact" toml:"contact" validate:"omitempty,email"`
	}

	// VersionManifest defines the version information in a module's manifest.
	VersionManifest struct {
		Version   string `json:"version" toml:"version" validate:"required"`
		SDKCompat string `json:"sdk_compat" toml:"sdk_compat"`
	}

	// Manifest defines a Cosmos SDK module manifest. It translates directly into
	// a Module model.
	Manifest struct {
		Module     ModuleManifest    `json:"module" toml:"module"`
		BugTracker BugTackerManifest `json:"bug_tracker" toml:"bug_tracker" validate:"omitempty,dive"`
		Authors    []AuthorsManifest `json:"authors" toml:"authors" validate:"required,gt=0,unique=Name,dive"`
		Version    VersionManifest   `json:"version" toml:"version" validate:"required,dive"`
	}
)

// ModuleFromManifest converts a Manifest to a Module model.
func ModuleFromManifest(manifest Manifest) models.Module {
	authors := make([]models.User, len(manifest.Authors))
	for i, a := range manifest.Authors {
		authors[i] = models.User{Name: a.Name, Email: models.NewNullString(a.Email)}
	}

	keywords := make([]models.Keyword, len(manifest.Module.Keywords))
	for i, k := range manifest.Module.Keywords {
		keywords[i] = models.Keyword{Name: k}
	}

	bugTracker := models.BugTracker{
		URL:     models.NewNullString(manifest.BugTracker.URL),
		Contact: models.NewNullString(manifest.BugTracker.Contact),
	}

	modVer := models.ModuleVersion{
		Version:   manifest.Version.Version,
		SDKCompat: models.NewNullString(manifest.Version.SDKCompat),
	}

	return models.Module{
		Name:          manifest.Module.Name,
		Team:          manifest.Module.Team,
		Repo:          manifest.Module.Repo,
		Description:   manifest.Module.Description,
		Documentation: manifest.Module.Documentation,
		Homepage:      manifest.Module.Homepage,
		Version:       modVer,
		Authors:       authors,
		Keywords:      keywords,
		BugTracker:    bugTracker,
	}
}
