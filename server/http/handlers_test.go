package http

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIndex(t *testing.T) {
	testcases := []struct {
		name      string
		viewData  []byte
		viewError error
	}{
		{"test no error", []byte("view data"), nil},
		{"test view error", nil, ErrInternal},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := new(mocks.View)
			rw := new(mocks.ResponseWriter)

			v.On("Index").
				Return(tc.viewData, tc.viewError)

			if tc.viewError == nil {
				v.On("ContentType").Return("text/plain")
				rw.On("Header").Return(http.Header{})
				rw.On("WriteHeader", http.StatusOK).Return()
				rw.On("Write", tc.viewData).Return(len(tc.viewData), nil)
			}

			h := newHandlers(nil, v)

			err := h.Index(rw, nil)

			if tc.viewError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.viewError))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSearchQuery(t *testing.T) {
	testcases := []struct {
		name       string
		query      string
		queryError error
	}{
		{"test no error", "foo", nil},
		{"test empty query", "", ErrBadRequest},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ws := new(mocks.WorkoutService)
			v := new(mocks.View)

			if tc.query != "" {
				ws.On("CreateWorkout", mock.Anything, tc.query).
					Return(&domain.Workout{}, nil)

				v.On("ComposeSearchData", mock.Anything, mock.Anything).
					Return([]byte("view data"), nil)

				v.On("ContentType").Return("text/plain")
			}

			h := newHandlers(ws, v)

			rw := new(mocks.ResponseWriter)
			if tc.query != "" {
				rw.On("Header").Return(http.Header{})
				rw.On("WriteHeader", http.StatusOK).Return()
				rw.On("Write", []byte("view data")).Return(len("view data"), nil)
			}

			r := buildSearchRequest(tc.query)

			err := h.Search(rw, r)

			if tc.queryError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.queryError))
			} else {
				assert.NoError(t, err)
			}

			ws.AssertExpectations(t)
			v.AssertExpectations(t)
			rw.AssertExpectations(t)
		})
	}
}

func TestSearchWorkoutService(t *testing.T) {
	testcases := []struct {
		name         string
		serviceData  *domain.Workout
		serviceError error
	}{
		{"test no error", &domain.Workout{Name: "test"}, nil},
		{"test empty query", nil, ErrInternal},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ws := new(mocks.WorkoutService)
			v := new(mocks.View)

			ws.On("CreateWorkout", mock.Anything, mock.Anything).
				Return(tc.serviceData, tc.serviceError)

			if tc.serviceError == nil {
				v.On("ComposeSearchData", mock.Anything, tc.serviceData).
					Return([]byte("view data"), nil)

				v.On("ContentType").Return("text/plain")
			}

			h := newHandlers(ws, v)

			rw := new(mocks.ResponseWriter)
			if tc.serviceError == nil {
				rw.On("Header").Return(http.Header{})
				rw.On("WriteHeader", http.StatusOK).Return()
				rw.On("Write", mock.Anything).Return(1, nil)
			}

			r := buildSearchRequest("foo")

			err := h.Search(rw, r)

			if tc.serviceError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.serviceError))
			} else {
				assert.NoError(t, err)
			}

			ws.AssertExpectations(t)
			v.AssertExpectations(t)
			rw.AssertExpectations(t)
		})
	}
}

func TestSearchView(t *testing.T) {
	testcases := []struct {
		name      string
		viewData  []byte
		viewError error
	}{
		{"test no error", []byte("view data"), nil},
		{"test empty query", nil, ErrInternal},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ws := new(mocks.WorkoutService)
			v := new(mocks.View)

			ws.On("CreateWorkout", mock.Anything, mock.Anything).
				Return(&domain.Workout{}, nil)

			v.On("ComposeSearchData", mock.Anything, mock.Anything).
				Return(tc.viewData, tc.viewError)

			if tc.viewError == nil {
				v.On("ContentType").Return("text/plain")
			}

			h := newHandlers(ws, v)

			rw := new(mocks.ResponseWriter)
			if tc.viewError == nil {
				rw.On("Header").Return(http.Header{})
				rw.On("WriteHeader", http.StatusOK).Return()
				rw.On("Write", tc.viewData).Return(len(tc.viewData), nil)
			}

			r := buildSearchRequest("foo")

			err := h.Search(rw, r)

			if tc.viewError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.viewError))
			} else {
				assert.NoError(t, err)
			}

			ws.AssertExpectations(t)
			v.AssertExpectations(t)
			rw.AssertExpectations(t)
		})
	}
}

func buildSearchRequest(query string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/search?q=%s", query), nil)
	return r
}
