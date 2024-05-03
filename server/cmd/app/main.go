package main

import (
	"os"

	"github.com/brendenehlers/workout-app/server/config"
	"github.com/brendenehlers/workout-app/server/html"
	"github.com/brendenehlers/workout-app/server/http"
	"github.com/brendenehlers/workout-app/server/log"
	"github.com/brendenehlers/workout-app/server/workout"
)

func main() {
	ws := workout.New()
	v := html.New(config.PagesDir)

	dev := os.Getenv(config.APP_ENV) == config.DEVELOPMENT
	addr := ":4321"

	server := http.New(ws, v, http.ServerConfig{
		Addr: addr,
		Dev:  dev,
	})
	log.Err(server.Start())
}
