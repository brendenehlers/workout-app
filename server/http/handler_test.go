package http

import (
	"errors"
	"net/http"
	"testing"

	"github.com/brendenehlers/workout-app/server/mocks"
)

func Test_handle(t *testing.T) {
	testCases := []struct {
		name   string
		writes []string
		status int
		err    *wrappedResponseError
	}{
		{"test no writes", []string{}, http.StatusOK, nil},
		{"test writes", []string{"test"}, http.StatusOK, nil},
		{"test error", []string{}, http.StatusInternalServerError, WrapError(errors.New("test"), ErrInternal)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fn := func(w http.ResponseWriter, r *http.Request) error {
				w.WriteHeader(tc.status)

				if tc.err != nil {
					return tc.err
				}

				for _, write := range tc.writes {
					w.Write([]byte(write))
				}

				return nil
			}

			rw := &mocks.ResponseWriter{}
			r := &http.Request{}

			rw.On("Header").Maybe().Return(http.Header{})

			if tc.err == nil {
				rw.On("WriteHeader", tc.status).
					Return()

				for _, write := range tc.writes {
					rw.On("Write", []byte(write)).
						Return(len(write), nil)
				}
			} else {
				rw.On("WriteHeader", tc.status).
					Return()

				rw.On("Write", []byte(tc.err.apiError.msg)).
					Return(len(tc.err.apiError.msg), nil)
			}

			handle(fn, rw, r)

			rw.AssertExpectations(t)
		})
	}
}
