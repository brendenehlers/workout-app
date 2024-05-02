package http

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	ws := new(mocks.WorkoutService)
	ws.On("CreateWorkout", "foo").
		Return(&domain.Workout{}, nil)

	v := new(mocks.View)
	v.On("ComposeSearchData", context.Background(), &domain.Workout{}).
		Return([]byte("test data"), nil)

	h := newHandlers(ws, v)

	rw := &responseWriter{}
	r := buildSearchRequest("foo")

	err := h.Search(rw, r)

	ws.AssertExpectations(t)
	v.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(rw.writes))
	assert.Equal(t, "test data", string(rw.writes[0]))
}

func TestSearchInvalidQuery(t *testing.T) {
	ws := new(mocks.WorkoutService)
	v := new(mocks.View)
	h := newHandlers(ws, v)

	rw := &responseWriter{}
	r := buildSearchRequest("")

	err := h.Search(rw, r)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrBadRequest))
}

func buildSearchRequest(query string) *http.Request {
	return &http.Request{
		URL: &url.URL{
			Path:     "/search",
			RawQuery: "q=" + query,
		},
	}
}
