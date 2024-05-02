package http

import (
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

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.Route("/", func(r chi.Router) {
		handlers := newHandlers(ws, v)
		r.Get("/", wrapHandler(handlers.Index))
		r.Get("/search", wrapHandler(handlers.Search))
	})

	return &Server{
		Server: &http.Server{
			Handler: r,
			Addr:    cfg.Addr,
		},
	}
}

func (s *Server) Start() error {
	log.Infof("Server listening on %s", s.Addr)
	return s.ListenAndServe()
}
