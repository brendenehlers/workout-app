package http

import (
	"net/http"

	"github.com/brendenehlers/workout-app/server/log"
)

func Logger(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
