package http

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapError(t *testing.T) {
	err := errors.New("test error")
	apiErr := &apiError{msg: "test error", status: 400}

	wrappedErr := WrapError(err, apiErr)

	assert.Equal(t, err, wrappedErr.error)
	assert.Equal(t, apiErr, wrappedErr.apiError)
}
