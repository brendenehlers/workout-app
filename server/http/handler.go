package http

import (
	"fmt"
	"net/http"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/log"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	writes [][]byte
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
}

func (rw *responseWriter) Write(bytes []byte) (int, error) {
	rw.writes = append(rw.writes, bytes)
	return len(bytes), nil
}

func (rw *responseWriter) flush() error {
	rw.ResponseWriter.WriteHeader(rw.status)

	for _, write := range rw.writes {
		_, err := rw.ResponseWriter.Write(write)
		if err != nil {
			return err
		}
	}

	return nil
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func (fn handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handle(fn, w, r)
}

func wrapHandler(fn handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handle(fn, w, r)
	}
}

func handle(fn handlerFunc, w http.ResponseWriter, r *http.Request) {
	rw := &responseWriter{ResponseWriter: w}

	if err := fn(rw, r); err != nil {
		handleError(w, err)
	}

	// past the point of no return
	err := rw.flush()
	if err != nil {
		log.Err(err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case domain.WrappedError:
		log.Errorf(e.Error())
		msg, status := e.APIError()
		writeError(w, msg, status)
	case error:
		log.Errorf(e.Error())
		writeError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	default:
		panic("invalid error type passed to handleError")
	}
}

// http.Error() adds a newline at the end of the string
// don't want that
func writeError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprint(w, error)
}
