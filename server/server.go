package server

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	router     *mux.Router
	controller Controller
}

func New(ctrl Controller) *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}
