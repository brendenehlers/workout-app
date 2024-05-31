package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/brendenehlers/workout-app/services/workout/internals/config"
	"github.com/brendenehlers/workout-app/services/workout/internals/mocks"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func setupTestServer(router chi.Router) (*Server, *httptest.Server) {
	ts := httptest.NewServer(router)
	s := &Server{
		router:     router,
		httpServer: ts.Config,
	}

	return s, ts
}

func TestMain(m *testing.M) {
	envs := os.Environ()
	os.Setenv(config.ENV_PORT, "8080")

	code := m.Run()

	os.Clearenv()
	for _, env := range envs {
		parts := strings.Split(env, "=")
		os.Setenv(parts[0], parts[1])
	}

	os.Exit(code)
}

func TestNewServerCreatesServer(t *testing.T) {
	server := NewServer(chi.NewRouter())

	assert.NotNil(t, server)
}

func TestNewServerAddsRouter(t *testing.T) {
	server := NewServer(chi.NewRouter())

	assert.NotNil(t, server.router)
}

func TestNewServerCreatesHttpServer(t *testing.T) {
	server := NewServer(chi.NewRouter())

	assert.NotNil(t, server.httpServer)
}

func TestNewServerAddsRouterToHttpServer(t *testing.T) {
	server := NewServer(chi.NewRouter())

	assert.NotNil(t, server.httpServer.(*http.Server).Handler)
}

func TestNewServerAddsPortToHttpServer(t *testing.T) {
	port := "8080"
	server := NewServer(chi.NewRouter())

	assert.Equal(t, conformPort(port), server.httpServer.(*http.Server).Addr)
}

func TestConformPortAddsSemiColon(t *testing.T) {
	port := conformPort("8080")

	assert.Equal(t, ":8080", port)
}

func TestConformPortDoesntAddSemiColonWhenPresent(t *testing.T) {
	port := conformPort(":8080")

	assert.Equal(t, ":8080", port)
}

func TestRegisterRoutesAddsRoutes(t *testing.T) {
	s := NewServer(chi.NewRouter())
	called := false

	mockRouter := mocks.MockRoutable{
		RegisterRoutesFn: func(r *chi.Router) {
			assert.Equal(t, &s.router, r)
			called = true
		},
	}
	err := s.RegisterRoutes([]Routable{&mockRouter})

	assert.NoError(t, err)
	if !called {
		assert.Fail(t, "RegisterRoutesFn was not called")
	}
}

func TestRegisterRoutesReturnsErrorWhenNoRoutesProvided(t *testing.T) {
	s := NewServer(chi.NewRouter())

	err := s.RegisterRoutes([]Routable{})

	assert.ErrorIs(t, err, ErrNoAPIsProvided)
}

func TestStartServerStartsServer(t *testing.T) {
	router := chi.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s, ts := setupTestServer(router)
	defer ts.Close()
	ctx := context.Background()
	go func() {
		if err := s.Start(ctx); err != nil {
			assert.FailNow(t, err.Error())
		}
	}()

	res, err := http.Get(ts.URL)
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestStartIsBlockingUntilDoneSignalReceived(t *testing.T) {
	s := NewServer(chi.NewRouter())
	ctx, cancel := context.WithCancel(context.Background())
	blocking := true
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.Start(ctx)
		blocking = false
	}()

	cancel()
	wg.Wait()
	assert.False(t, blocking)
}

func TestStartReturnsListenAndServeError(t *testing.T) {
	s := &Server{
		router: chi.NewRouter(),
		httpServer: &mocks.MockServer{
			ListenAndServeFunc: func() error { return ErrListenAndServe },
			Server:             &http.Server{},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	err := s.Start(ctx)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrListenAndServe)
}

func TestStartDoesNotReturnErrServerClosed(t *testing.T) {
	s := &Server{
		router: chi.NewRouter(),
		httpServer: &mocks.MockServer{
			ListenAndServeFunc: func() error { return http.ErrServerClosed },
			Server:             &http.Server{},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	err := s.Start(ctx)

	assert.Nil(t, err)
}

func TestStartShutsDownGracefullyWhenDone(t *testing.T) {
	shutdownCalled := false
	var wg sync.WaitGroup
	wg.Add(1)
	s := &Server{
		router: chi.NewRouter(),
		httpServer: &mocks.MockServer{
			ShutdownFunc: func(_ context.Context) error {
				shutdownCalled = true
				defer wg.Done()
				return nil
			},
			Server: &http.Server{},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	var err error
	go func() {
		err = s.Start(ctx)
	}()

	cancel()
	wg.Wait()
	assert.True(t, shutdownCalled)
	assert.Nil(t, err)
}

func TestStartShutdownErrorIsReturned(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	s := &Server{
		router: chi.NewRouter(),
		httpServer: &mocks.MockServer{
			ShutdownFunc: func(_ context.Context) error {
				wg.Done()
				return ErrShutdownError
			},
			Server: &http.Server{},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	var err error
	go func() {
		err = s.Start(ctx)
	}()

	cancel()
	wg.Wait()
	assert.ErrorIs(t, err, ErrShutdownError)
}

func TestStartShutdownTimesOutAfterFiveSeconds(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	s := &Server{
		router: chi.NewRouter(),
		httpServer: &mocks.MockServer{
			ShutdownFunc: func(ctx context.Context) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(5 * time.Second):
					return nil
				}
			},
			Server: &http.Server{},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	var err error
	go func() {
		err = s.Start(ctx)
		wg.Done()
	}()

	cancel()
	wg.Wait()
	assert.ErrorIs(t, err, ErrShutdownTimeout)
}
