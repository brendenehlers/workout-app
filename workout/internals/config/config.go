package config

import (
	"os"

	"github.com/brendenehlers/workout-app/services/workout/log"
	"github.com/joho/godotenv"
)

const (
	ENV_PORT = "PORT"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Error(err.Error(), nil)
		panic("Failed to load env files")
	}
}

func Env(key string) string {
	return os.Getenv(key)
}
