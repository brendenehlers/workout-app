package html

import (
	"bytes"
	"context"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/html/templates"
)

type HTMLView struct{}

var _ domain.View = HTMLView{}

func New() *HTMLView {
	return &HTMLView{}
}

func (HTMLView) ComposeSearchData(ctx context.Context, w *domain.Workout) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := templates.Workout(w).Render(ctx, buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
