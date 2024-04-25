package http

import (
	"net/http"

	"github.com/brendenehlers/workout-app/server/domain"
)

type handlers struct {
	view    domain.View
	service domain.WorkoutService
}

func (h *handlers) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var body domain.CreateWorkoutRequest
	err := readJSON(r.Body, &body)
	if err != nil {
		panic(err)
	}

	workout, err := h.service.CreateWorkout(body)
	if err != nil {
		panic(err)
	}

	err = h.view.EncodeContent(w, workout)
	if err != nil {
		panic(err)
	}
}
