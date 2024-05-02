package html

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/brendenehlers/workout-app/server/domain"
	"github.com/brendenehlers/workout-app/server/html/templates"
)

type HTMLView struct{}

var _ domain.View = HTMLView{}

func New() *HTMLView {
	return &HTMLView{}
}

func (HTMLView) ContentType() string {
	return "text/html"
}

func (HTMLView) Index() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	file, err := os.Open("public/pages/index.html")
	if err != nil {
		return nil, err
	}
	defer file.Close()

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
