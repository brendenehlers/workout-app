package domain

import "context"

type View interface {
	ComposeSearchData(context.Context, *Workout) ([]byte, error)
}
