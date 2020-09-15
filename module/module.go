package module

// Keyword defines a module keyword, where a module can have one or more keywords.
type Keyword struct {
	ID   int    `json:"-" yaml:"-" db:"id"`
	Name string `json:"name" yaml:"name" db:"name"`
}

type (
	// Author defines a type alias for a User
	Author User

	// Contributors defines a list of User types
	Contributors []User

	// User defines an entity that contributes to a Module type.
	User struct {
		ID                int    `json:"-" yaml:"-" db:"id"`
		Name              string `json:"name" yaml:"name" db:"name"`
		URL               string `json:"url" yaml:"url" db:"url"`
		Email             string `json:"email" yaml:"email" db:"email"`
		GithubAccessToken string `json:"github_access_token" yaml:"github_access_token" db:"github_access_token"`
		APIToken          string `json:"api_token" yaml:"api_token" db:"api_token"`
	}
)

// Bug defines the metadata information for reporting bug reports on a given
// Module type.
type Bug struct {
	ID      int    `json:"-" yaml:"-" db:"id"`
	URL     string `json:"url" yaml:"url" db:"url"`
	Contact string `json:"contact" yaml:"contact" db:"contact"`
}

// Module defines a Cosmos SDK module.
type Module struct {
	ID          int    `json:"-" yaml:"-" db:"id"`
	Name        string `json:"name" yaml:"name" db:"name"`
	Description string `json:"description" yaml:"description" db:"description"`
	Version     string `json:"version" yaml:"version" db:"version"`
	Homepage    string `json:"homepage" yaml:"homepage" db:"homepage"`
	Repo        string `json:"repo" yaml:"repo" db:"repo"`
	BugID       int    `json:"-" yaml:"-" db:"bug_id"`
	Author      int    `json:"-" yaml:"-" db:"author"`
}
