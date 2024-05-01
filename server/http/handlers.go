package http

import (
	"net/http"
	"path/filepath"

	"github.com/brendenehlers/workout-app/server/domain"
)

type handlers struct {
	ws domain.WorkoutService
	v  domain.View
}

func newHandlers(ws domain.WorkoutService, v domain.View) *handlers {
	return &handlers{
		ws: ws,
		v:  v,
	}
}

func (h *handlers) Index(w http.ResponseWriter, r *http.Request) error {
	fp := filepath.Join("public", "pages", "index.html")
	http.ServeFile(w, r, fp)

	return nil
}

func (h *handlers) Search(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query().Get("q")

	if query == "" {
		return WrapError(nil, ErrBadRequest)
	}

	workout, err := h.ws.CreateWorkout(query)
	if err != nil {
		return err
	}

	data, err := h.v.ComposeSearchData(r.Context(), workout)
	if err != nil {
		return err
	}

	w.Write(data)

	return nil
}
