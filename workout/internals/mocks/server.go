package mocks

import (
	"context"
	"fmt"
	"net/http"
)

type MockServer struct {
	ListenAndServeFunc func() error
	ShutdownFunc       func(context.Context) error
	*http.Server
}

func (m *MockServer) ListenAndServe() error {
	if m.ListenAndServeFunc == nil {
		fmt.Println("no ListenAndServeFunc")
		return nil
	}
	return m.ListenAndServeFunc()
}

func (m *MockServer) Shutdown(ctx context.Context) error {
	if m.ShutdownFunc == nil {
		fmt.Println("no ShutdownFunc")
		return nil
	}
	return m.ShutdownFunc(ctx)
}
