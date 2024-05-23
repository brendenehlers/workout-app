package server

import (
	"os"
	"strings"
	"testing"

	"github.com/brendenehlers/workout-app/services/workout/internals/config"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func loadEnv(t *testing.T, vars map[string]string) {
	os.Chdir(t.TempDir())
	godotenv.Write(vars, ".env")
	config.Init()
}

func TestMain(m *testing.M) {
	envs := os.Environ()

	code := m.Run()

	os.Clearenv()
	for _, env := range envs {
		parts := strings.Split(env, "=")
		os.Setenv(parts[0], parts[1])
	}

	os.Exit(code)
}

func TestNewServerCreatesServer(t *testing.T) {
	port := "8080"
	loadEnv(t, map[string]string{
		config.ENV_PORT: port,
	})

	server := NewServer(chi.NewRouter())

	assert.NotNil(t, server)
}

func TestNewServerAddsRouter(t *testing.T) {
	port := "ADDS ROUTER PORT 8080"
	loadEnv(t, map[string]string{
		config.ENV_PORT: port,
	})

	server := NewServer(chi.NewRouter())

	assert.NotNil(t, server.router)
}

func TestNewServerCreatesHttpServer(t *testing.T) {
	port := "8080"
	loadEnv(t, map[string]string{
		config.ENV_PORT: port,
	})

	server := NewServer(chi.NewRouter())

	assert.NotNil(t, server.httpServer)
}

func TestNewServerAddsRouterToHttpServer(t *testing.T) {
	port := "8080"
	loadEnv(t, map[string]string{
		config.ENV_PORT: port,
	})

	server := NewServer(chi.NewRouter())

	assert.NotNil(t, server.httpServer.Handler)
}

func TestNewServerAddsPortToHttpServer(t *testing.T) {
	port := "8080"
	loadEnv(t, map[string]string{
		config.ENV_PORT: port,
	})

	server := NewServer(chi.NewRouter())

	assert.Equal(t, conformPort(port), server.httpServer.Addr)
}

func TestConformPortAddsSemiColon(t *testing.T) {
	port := conformPort("8080")

	assert.Equal(t, ":8080", port)
}

func TestConformPortDoesntAddSemiColonWhenPresent(t *testing.T) {
	port := conformPort(":8080")

	assert.Equal(t, ":8080", port)
}
