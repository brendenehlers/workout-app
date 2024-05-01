package mocks

import (
	"context"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/stretchr/testify/mock"
)

type MockView struct {
	mock.Mock
}

var _ domain.View = &MockView{}

func (m *MockView) ComposeSearchData(ctx context.Context, w *domain.Workout) ([]byte, error) {
	args := m.Called(ctx, w)
	return args.Get(0).([]byte), args.Error(1)
}
