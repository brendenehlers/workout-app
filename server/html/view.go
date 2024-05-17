package html

import (
	"bytes"
	"context"

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

func (HTMLView) Error(ctx context.Context, msg string) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := templates.Error(msg).Render(ctx, buf)
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
