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

func viewWrapHandler(v domain.View, fn handlerFunc) http.HandlerFunc {
	vw := viewWrapper{view: v, fn: fn}
	return func(w http.ResponseWriter, r *http.Request) {
		vw.handle(w, r)
	}
}

type viewWrapper struct {
	view domain.View
	fn   handlerFunc
}

func (vw viewWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vw.handle(w, r)
}

func (vw viewWrapper) handle(w http.ResponseWriter, r *http.Request) {
	rw := &responseWriter{ResponseWriter: w}

	if err := vw.fn(rw, r); err != nil {
		vw.handleError(w, r, err)
	}

	// past the point of no return
	err := rw.flush()
	if err != nil {
		log.Err(err)
	}
}

func (vw viewWrapper) handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Err(err)
	switch e := err.(type) {
	case domain.WrappedError:
		msg, status := e.APIError()
		vw.writeError(w, r, msg, status)
	case error:
		vw.writeError(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	default:
		panic("invalid error type passed to handleError")
	}
}

func (vw viewWrapper) writeError(w http.ResponseWriter, r *http.Request, msg string, code int) {
	data, err := vw.view.Error(r.Context(), msg)
	if err != nil {
		log.Err(err)
		writeDefaultError(w)
		return
	}

	w.Header().Set("Content-Type", vw.view.ContentType())
	w.Write(data)
	w.WriteHeader(code)
}

// http.Error() adds a newline at the end of the string
// don't want that
func writeDefaultError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, http.StatusText(http.StatusInternalServerError))
}
