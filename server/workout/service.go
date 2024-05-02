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
	return &domain.Workout{
		Name:        "Killer Workout 3000",
		Description: "Gonna hurt",
		Exercises: []domain.Exercise{
			{
				Name:        "Pushups",
				Description: "Do pushups",
			},
			{
				Name:        "Situps",
				Description: "Do situps",
			},
		},
	}, nil
}
