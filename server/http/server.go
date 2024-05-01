package http

import (
	"net/http"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/log"
)

type Server struct {
	*http.Server
}

func New(addr string, service domain.WorkoutService) *Server {
	mux := http.NewServeMux()
	handlers := &handlers{
		service: service,
	}

	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	mux.HandleFunc("GET /", handlers.Index)
	mux.HandleFunc("GET /search", handlers.Search)

	return &Server{
		Server: &http.Server{
			Handler: Logger(mux),
			Addr:    addr,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Server listening on %s\n", s.Addr)
	return s.ListenAndServe()
}
