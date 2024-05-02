package domain

import "context"

type View interface {
	ContentType() string

	Index() ([]byte, error)
	ComposeSearchData(context.Context, *Workout) ([]byte, error)
}
