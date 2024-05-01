package main

import (
	"os"

	"github.com/brendenehlers/workout-app/server/config"
	"github.com/brendenehlers/workout-app/server/http"
	"github.com/brendenehlers/workout-app/server/log"
)

func main() {
	dev := os.Getenv(config.APP_ENV) == config.DEVELOPMENT
	addr := ":8080"

	server := http.New(http.ServerConfig{
		Addr: addr,
		Dev:  dev,
	})
	log.Fatal(server.Start())
}
