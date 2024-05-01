package http

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/http/templates"
)

type handlers struct {
	service domain.WorkoutService
}

func (h *handlers) Index(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("public", "pages", "index.html")
	http.ServeFile(w, r, fp)
}

func (h *handlers) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	fmt.Println(query)

	templ.Handler(templates.SearchQuery(query)).ServeHTTP(w, r)
}
