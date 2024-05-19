package http

import (
	"context"
	"net/http"

	"github.com/brendenehlers/workout-app/server/log"
	"github.com/google/uuid"
)

type keyType string

const TraceIDKey keyType = "traceId"

func traceId(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		traceId := uuid.New().String()

		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		ctx = context.WithValue(ctx, TraceIDKey, traceId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(f)
}

func logger(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w}

		defer func() {
			log.Debug("http request", log.Fields{
				"remoteAddr":  r.RemoteAddr,
				"contextPath": r.URL.Path,
				"query":       r.URL.Query().Encode(),
				"userAgent":   r.UserAgent(),
				"method":      r.Method,
				"status":      rw.status,
				"traceId":     r.Context().Value(TraceIDKey),
			})

		}()

		next.ServeHTTP(rw, r)
		rw.flush()
	}

	return http.HandlerFunc(f)
}
