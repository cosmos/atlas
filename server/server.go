package server

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

type Server struct {
	router     *mux.Router
	controller Controller
}

func New(ctrl Controller) *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

// func openDatabase(cfg Config) (*sqlx.DB, error) {
// 	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s")
// 	return sqlx.Connect("postgres", connStr)
// }
