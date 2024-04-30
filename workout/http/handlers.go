package http

import (
	"net/http"

	"github.com/brendenehlers/workout-app/services/workout/domain"
)

type handlers struct {
	s domain.SearchService
}

func (h handlers) createNewWorkout(w http.ResponseWriter, r *http.Request) {
	var query domain.WorkoutQuery
	err := readJSON(r.Body, &query)
	if err != nil {
		panic(r)
	}

	workout, err := h.s.Search(query)
	if err != nil {
		panic(err)
	}

	writeJSON(w, workout)
}
