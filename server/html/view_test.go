package html

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	view := HTMLView{}

	data, err := view.Error(context.Background(), "test")

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Contains(t, string(data), "test")
}
