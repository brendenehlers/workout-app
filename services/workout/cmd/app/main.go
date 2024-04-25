package main

import (
	"github.com/brendenehlers/workout-app/services/workout/domain"
	"github.com/brendenehlers/workout-app/services/workout/http"
	"github.com/brendenehlers/workout-app/services/workout/log"
)

func main() {
	s := Searcher{}
	server := http.New(":8080", &s)
	log.Fatal(server.Start())
}

type Searcher struct{}

func (*Searcher) Search(q domain.WorkoutQuery) (*domain.WorkoutData, error) {
	log.Printf("Search called on query: %s", q.Query)

	return &domain.WorkoutData{
		Id: q.Query,
	}, nil
}
