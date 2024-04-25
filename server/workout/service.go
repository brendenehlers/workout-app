package workout

import "github.com/brendenehlers/workout-app/server/domain"

type WorkoutService struct{}

var _ domain.WorkoutService = WorkoutService{}

func New() *WorkoutService {
	return &WorkoutService{}
}

func (WorkoutService) CreateWorkout(request domain.CreateWorkoutRequest) (*domain.Workout, error) {
	return &domain.Workout{
		Name:        "My new workout",
		Description: "Will destroy you",
	}, nil
}
