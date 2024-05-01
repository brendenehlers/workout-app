package main

import (
	"github.com/brendenehlers/workout-app/server/http"
	"github.com/brendenehlers/workout-app/server/log"
)

func main() {
	server := http.New(":8080")
	log.Fatal(server.Start())
}
