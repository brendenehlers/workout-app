package http

import (
	"net/http"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/log"
)

type Server struct {
	*http.Server
}

func New(addr string, view domain.View, service domain.WorkoutService) *Server {
	mux := http.NewServeMux()
	handlers := &handlers{
		view:    view,
		service: service,
	}

	mux.HandleFunc("POST /workout/create", handlers.CreateWorkout)

	return &Server{
		Server: &http.Server{
			Handler: mux,
			Addr:    addr,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Server listening on %s\n", s.Addr)
	return s.ListenAndServe()
}
