package http

import (
	"fmt"
	"net/http"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/log"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	*http.Server
}

type ServerConfig struct {
	Addr string
	Dev  bool
}

func New(ws domain.WorkoutService, v domain.View, cfg ServerConfig) *Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, traceId)

	if cfg.Dev {
		r.Use(middleware.NoCache)
	}

	r.Route("/api", func(r chi.Router) {
		handlers := newHandlers(ws, v)
		r.Get("/search", viewWrapHandler(v, handlers.Search))
	})

	return &Server{
		Server: &http.Server{
			Handler: r,
			Addr:    cfg.Addr,
		},
	}
}

func (s *Server) Start() error {
	log.Info(fmt.Sprintf("Server listening on %s", s.Addr))
	return s.ListenAndServe()
}
