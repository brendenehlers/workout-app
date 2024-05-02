package http

import "net/http"

type WrappedError interface {
	APIError() (string, int)
	error
}

var (
	ErrBadRequest = &apiError{msg: "invalid input", status: http.StatusBadRequest}
	ErrInternal   = &apiError{msg: "internal server error", status: http.StatusInternalServerError}
)

func WrapError(err error, apiErr *apiError) *wrappedResponseError {
	return &wrappedResponseError{error: err, apiError: apiErr}
}

type wrappedResponseError struct {
	error
	apiError *apiError
}

func (e wrappedResponseError) Is(err error) bool {
	return e.apiError == err
}

func (e wrappedResponseError) APIError() (string, int) {
	return e.apiError.APIError()
}

type apiError struct {
	msg    string
	status int
}

func (e apiError) APIError() (string, int) {
	return e.msg, e.status
}

func (e apiError) Error() string {
	return e.msg
}
