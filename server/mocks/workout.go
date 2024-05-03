package mocks

import (
	"context"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/stretchr/testify/mock"
)

type WorkoutService struct {
	mock.Mock
}

var _ domain.WorkoutService = &WorkoutService{}

func (s *WorkoutService) CreateWorkout(ctx context.Context, query string) (*domain.Workout, error) {
	args := s.Called(ctx, query)
	return args.Get(0).(*domain.Workout), args.Error(1)
}
