package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Create new Router for web-server
func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func (s *Server) ConfigureRouter() {
	// Routes
	s.Router.HandleFunc("/", s.Index)

	// Configure static files
	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
