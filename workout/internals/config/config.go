package config

import (
	"fmt"
	"os"

	"github.com/brendenehlers/workout-app/services/workout/log"
	"github.com/joho/godotenv"
)

const (
	ENV_PORT = "PORT"
)

var (
	ErrNotFound = func(key string) error { return fmt.Errorf("key '%s' not found", key) }
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Error(err.Error(), nil)
		panic("Failed to load env files")
	}
}

func Env(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(ErrNotFound(key))
	}
	return val
}
