package mocks

import (
	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/stretchr/testify/mock"
)

type MockWorkoutService struct {
	mock.Mock
}

var _ domain.WorkoutService = &MockWorkoutService{}

func (s *MockWorkoutService) CreateWorkout(query string) (*domain.Workout, error) {
	args := s.Called(query)
	return args.Get(0).(*domain.Workout), args.Error(1)
}
