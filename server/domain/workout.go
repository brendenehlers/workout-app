package domain

import "context"

type Exercise struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type Workout struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Exercises   []Exercise `json:"exercises"`
}

type WorkoutService interface {
	CreateWorkout(context.Context, string) (*Workout, error)
}
