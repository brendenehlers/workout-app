package domain

import (
	"net/http"
)

type View interface {
	EncodeContent(http.ResponseWriter, any) error
}
