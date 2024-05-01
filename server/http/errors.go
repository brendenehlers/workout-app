package http

type APIError interface {
	APIError() (string, int)
}

var (
	ErrBadRequest = &apiError{msg: "invalid input", status: 400}
)

func WrapError(err error, apiErr *apiError) error {
	return wrappedResponseError{error: err, apiError: apiErr}
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
