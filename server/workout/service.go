package workout

import (
	"github.com/brendenehlers/workout-app/server/domain"
)

type WorkoutService struct{}

var _ domain.WorkoutService = WorkoutService{}

func New() *WorkoutService {
	return &WorkoutService{}
}

func (WorkoutService) CreateWorkout(query string) (*domain.Workout, error) {
	panic("not implemented")
}
