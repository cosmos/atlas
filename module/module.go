package module

type Keyword struct {
	ID   int    `json:"-" yaml:"-" db:"id"`
	Name string `json:"name" yaml:"name" db:"name"`
}

type (
	Author       User
	Contributors []User

	User struct {
		ID                int    `json:"-" yaml:"-" db:"id"`
		Name              string `json:"name" yaml:"name" db:"name"`
		URL               string `json:"url" yaml:"url" db:"url"`
		Email             string `json:"email" yaml:"email" db:"email"`
		GithubAccessToken string `json:"github_access_token" yaml:"github_access_token" db:"github_access_token"`
		APIToken          string `json:"api_token" yaml:"api_token" db:"api_token"`
	}
)

type Bug struct {
	ID      int    `json:"-" yaml:"-" db:"id"`
	URL     string `json:"url" yaml:"url" db:"url"`
	Contact string `json:"contact" yaml:"contact" db:"contact"`
}

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
