package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_traceId(t *testing.T) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		traceId := r.Context().Value(TraceIDKey)
		assert.NotNil(t, traceId)
	}

	handler := traceId(http.HandlerFunc(fn))

	r, _ := http.NewRequest("GET", "/", nil)

	traceId(handler).ServeHTTP(nil, r)
}
