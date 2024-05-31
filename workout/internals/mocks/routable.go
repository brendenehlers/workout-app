package mocks

import "github.com/go-chi/chi"

type MockRoutable struct {
	RegisterRoutesFn func(*chi.Router)
}

func (r *MockRoutable) RegisterRoutes(router *chi.Router) {
	r.RegisterRoutesFn(router)
}
