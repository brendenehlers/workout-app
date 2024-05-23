package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/brendenehlers/workout-app/services/workout/internals/config"
	"github.com/go-chi/chi"
)

type Routable interface {
	RegisterRoutes(router *chi.Router)
}

type Server struct {
	httpServer *http.Server
	router     chi.Router
}

func NewServer(router chi.Router) *Server {
	port := config.Env(config.ENV_PORT)

	return &Server{
		httpServer: &http.Server{
			Handler: router,
			Addr:    conformPort(port),
		},
		router: router,
	}
}

func conformPort(port string) string {
	if strings.HasPrefix(port, ":") {
		return port
	}
	return fmt.Sprintf(":%s", port)
}

func (s *Server) Start(ctx context.Context) error {
	panic("not implemented")
}

func (s *Server) Shutdown(ctx context.Context) error {
	panic("not implemented")
}

func (s *Server) RegisterRoutes(apis []Routable) {
	panic("not implemented")
}
