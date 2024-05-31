package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/brendenehlers/workout-app/services/workout/internals/config"
	"github.com/go-chi/chi"
)

var (
	ErrNoAPIsProvided  = errors.New("no apis provided")
	ErrListenAndServe  = errors.New("ListenAndServe error")
	ErrShutdownError   = errors.New("shutdown error")
	ErrShutdownTimeout = errors.New("shutdown timed out")
)

type Routable interface {
	RegisterRoutes(*chi.Router)
}

type ServerInterface interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

type Server struct {
	httpServer ServerInterface
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
	errCh := make(chan error)
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- errors.Join(ErrListenAndServe, err)
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		break
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return errors.Join(ErrShutdownTimeout, err)
		}

		return errors.Join(ErrShutdownError, err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	panic("not implemented")
}

func (s *Server) RegisterRoutes(apis []Routable) error {
	if len(apis) == 0 {
		return ErrNoAPIsProvided
	}

	for _, api := range apis {
		api.RegisterRoutes(&s.router)
	}
	return nil
}
