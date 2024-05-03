package domain

type WrappedError interface {
	APIError() (string, int)
	error
}
