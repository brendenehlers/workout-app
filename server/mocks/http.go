package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

var _ http.ResponseWriter = &ResponseWriter{}

type ResponseWriter struct {
	mock.Mock
}

func (m *ResponseWriter) WriteHeader(status int) {
	m.Called(status)
}

func (m *ResponseWriter) Write(bytes []byte) (int, error) {
	args := m.Called(bytes)
	return args.Int(0), args.Error(1)
}

func (m *ResponseWriter) Header() http.Header {
	args := m.Called()
	return args.Get(0).(http.Header)
}
