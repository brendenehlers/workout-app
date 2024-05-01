package http

import (
	"net/http"
	"os"

	"github.com/brendenehlers/workout-app/server/config"
	"github.com/brendenehlers/workout-app/server/log"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	*http.Server
}

func New(addr string) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	if os.Getenv(config.APP_ENV) == config.DEVELOPMENT {
		r.Use(middleware.NoCache)
	}

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	handlers := &handlers{}

	r.Get("/", handlers.Index)
	r.Get("/search", handlers.Search)

	return &Server{
		Server: &http.Server{
			Handler: r,
			Addr:    addr,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Server listening on %s\n", s.Addr)
	return s.ListenAndServe()
}
