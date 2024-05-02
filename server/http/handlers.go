package http

import (
	"net/http"

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
	data, err := h.v.Index()
	if err != nil {
		return WrapError(err, ErrInternal)
	}

	h.writeData(w, data)

	return nil
}

func (h *handlers) Search(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query().Get("q")

	if query == "" {
		return WrapError(nil, ErrBadRequest)
	}

	workout, err := h.ws.CreateWorkout(query)
	if err != nil {
		return WrapError(err, ErrInternal)
	}

	data, err := h.v.ComposeSearchData(r.Context(), workout)
	if err != nil {
		return WrapError(err, ErrInternal)
	}

	h.writeData(w, data)

	return nil
}

func (h *handlers) writeData(w http.ResponseWriter, data []byte) {
	w.Header().Add("Content-Type", h.v.ContentType())
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
