package main

import (
	"github.com/brendenehlers/workout-app/server/html"
	"github.com/brendenehlers/workout-app/server/http"
	"github.com/brendenehlers/workout-app/server/log"
	"github.com/brendenehlers/workout-app/server/workout"
)

func main() {
	view, service := html.New(), workout.New()
	server := http.New(":8080", view, service)
	log.Fatal(server.Start())
}
