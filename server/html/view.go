package html

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/html/templates"
)

type HTMLView struct {
	pagesDir string
}

var _ domain.View = HTMLView{}

func New(pagesDir string) *HTMLView {
	return &HTMLView{
		pagesDir: pagesDir,
	}
}

func (HTMLView) ContentType() string {
	return "text/html"
}

func (h HTMLView) Index() ([]byte, error) {
	file, err := os.Open(fmt.Sprintf("%s/index.html", h.pagesDir))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (HTMLView) ComposeSearchData(ctx context.Context, w *domain.Workout) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := templates.Workout(w).Render(ctx, buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
