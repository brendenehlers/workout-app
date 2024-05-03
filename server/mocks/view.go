package mocks

import (
	"context"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/stretchr/testify/mock"
)

type View struct {
	mock.Mock
}

var _ domain.View = &View{}

func (m *View) ContentType() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *View) Index() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (m *View) Error(ctx context.Context, msg string) ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (m *View) ComposeSearchData(ctx context.Context, w *domain.Workout) ([]byte, error) {
	args := m.Called(ctx, w)
	return args.Get(0).([]byte), args.Error(1)
}
