package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type keyType string

const TraceIDKey keyType = "traceId"

func traceId(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		traceId := uuid.New().String()

		ctx := r.Context()
		ctx = context.WithValue(ctx, TraceIDKey, traceId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(f)
}
