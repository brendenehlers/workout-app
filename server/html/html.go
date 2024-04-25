package html

import (
	"net/http"

	"github.com/brendenehlers/workout-app/server/domain"
)

type HTMLView struct{}

var _ domain.View = HTMLView{}

func New() *HTMLView {
	return &HTMLView{}
}

func (HTMLView) EncodeContent(w http.ResponseWriter, data any) error {
	w.Header().Add("Content-Type", "text/html")
	_, err := w.Write([]byte("<p>Hello, world</p>"))
	return err
}
