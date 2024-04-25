package main

import (
	"github.com/brendenehlers/workout-app/services/workout/embedding"
	"github.com/brendenehlers/workout-app/services/workout/http"
	"github.com/brendenehlers/workout-app/services/workout/log"
	"github.com/brendenehlers/workout-app/services/workout/search"
	"github.com/brendenehlers/workout-app/services/workout/vectorstore"
)

func main() {
	e := embedding.New()
	vs := vectorstore.New()
	s := search.New(e, vs)
	server := http.New(":8080", s)
	log.Fatal(server.Start())
}
