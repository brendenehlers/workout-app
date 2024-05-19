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
	r.Use(middleware.RequestID, logger, middleware.Recoverer, traceId)

	if cfg.Dev {
		r.Use(middleware.NoCache)
	}

	r.Route("/", func(r chi.Router) {
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
	log.Debug(fmt.Sprintf("Server listening on %s", s.Addr), nil)
	return s.ListenAndServe()
}
