package http

import (
	"errors"
	"net/http"
	"testing"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
			rw := new(mocks.ResponseWriter)
			r, _ := http.NewRequest(http.MethodGet, "/", nil)
			v := new(mocks.View)

			rw.On("Header").Maybe().Return(http.Header{})

			if tc.err == nil {
				rw.On("WriteHeader", tc.status).
					Return()

				for _, write := range tc.writes {
					rw.On("Write", []byte(write)).
						Return(len(write), nil)
				}
			} else {
				v.On("Error", mock.Anything, tc.err.apiError.msg).
					Return([]byte(tc.err.apiError.msg), nil)
				v.On("ContentType").Return("text/plain")

				rw.On("WriteHeader", tc.status).
					Return()
				rw.On("Write", []byte(tc.err.apiError.msg)).
					Return(len(tc.err.apiError.msg), nil)
			}

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

			vw := viewWrapper{
				view: v,
				fn:   fn,
			}

			vw.handle(rw, r)

			rw.AssertExpectations(t)
			v.AssertExpectations(t)
		})
	}
}

func Test_handleError(t *testing.T) {
	testCases := []struct {
		name string
		err  error
	}{
		{"test wrapped error", WrapError(errors.New("test"), ErrInternal)},
		{"test error", errors.New("test")},
		{"test unknown error type", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v := new(mocks.View)
			rw := new(mocks.ResponseWriter)
			r, _ := http.NewRequest(http.MethodGet, "/", nil)

			v.On("ContentType").Return("text/plain").Maybe()
			rw.On("Header").Return(http.Header{}).Maybe()
			rw.On("WriteHeader", mock.Anything).Return().Maybe()
			rw.On("Write", mock.Anything).Return(1, nil).Maybe()

			switch e := tc.err.(type) {
			case domain.WrappedError:
				msg, _ := e.APIError()
				v.On("Error", mock.Anything, msg).
					Return([]byte(msg), nil)
			case error:
				v.On("Error", mock.Anything, http.StatusText(http.StatusInternalServerError)).
					Return([]byte(http.StatusText(http.StatusInternalServerError)), nil)
			}

			vw := viewWrapper{
				view: v,
				fn:   nil,
			}

			if tc.err != nil {
				vw.handleError(rw, r, tc.err)
			} else {
				assert.Panics(t, func() {
					vw.handleError(rw, r, tc.err)
				})
			}

			rw.AssertExpectations(t)
			v.AssertExpectations(t)
		})
	}
}

func Test_writeError(t *testing.T) {
	testcases := []struct {
		name    string
		msg     string
		status  int
		content string
		err     error
	}{
		{"test success", "test", http.StatusOK, "test content", nil},
		{"test view error", "test", 0, "", errors.New("test")},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := new(mocks.View)
			rw := new(mocks.ResponseWriter)
			r, _ := http.NewRequest(http.MethodGet, "/", nil)

			v.On("Error", mock.Anything, tc.msg).
				Return([]byte(tc.content), tc.err)

			v.On("ContentType").Maybe().Return("text/plain")

			rw.On("Header").Maybe().Return(http.Header{})

			if tc.err != nil {
				rw.On("WriteHeader", http.StatusInternalServerError).
					Return()
				rw.On("Write", []byte(http.StatusText(http.StatusInternalServerError))).
					Return(len(http.StatusText(http.StatusInternalServerError)), tc.err)
			} else {
				rw.On("WriteHeader", tc.status).
					Return()
				rw.On("Write", []byte(tc.content)).
					Return(len(tc.content), nil)
			}

			vw := viewWrapper{
				view: v,
				fn:   nil,
			}

			vw.writeError(rw, r, tc.msg, tc.status)

			rw.AssertExpectations(t)
			v.AssertExpectations(t)
		})
	}
}
