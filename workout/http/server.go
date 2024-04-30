package http

import (
	"net/http"

	"github.com/brendenehlers/workout-app/services/workout/domain"
	"github.com/brendenehlers/workout-app/services/workout/log"
)

type Server struct {
	*http.Server
}

func New(addr string, search domain.SearchService) *Server {
	handler := http.NewServeMux()
	handlers := handlers{
		s: search,
	}

	handler.HandleFunc("POST /workout/new", handlers.createNewWorkout)

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on %s", s.Addr)
	return s.ListenAndServe()
}
