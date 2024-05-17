package domain

import (
	"context"
)

type View interface {
	ContentType() string

	Error(context.Context, string) ([]byte, error)
	ComposeSearchData(context.Context, *Workout) ([]byte, error)
}
