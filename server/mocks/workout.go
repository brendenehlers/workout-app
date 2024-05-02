package mocks

import (
	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/stretchr/testify/mock"
)

type WorkoutService struct {
	mock.Mock
}

var _ domain.WorkoutService = &WorkoutService{}

func (s *WorkoutService) CreateWorkout(query string) (*domain.Workout, error) {
	args := s.Called(query)
	return args.Get(0).(*domain.Workout), args.Error(1)
}
