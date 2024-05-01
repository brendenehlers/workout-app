package main

import (
	"github.com/brendenehlers/workout-app/server/http"
	"github.com/brendenehlers/workout-app/server/log"
	"github.com/brendenehlers/workout-app/server/workout"
)

func main() {
	service := workout.New()
	server := http.New(":8080", service)
	log.Fatal(server.Start())
}
