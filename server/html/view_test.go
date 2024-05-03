package html

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	testcases := []struct {
		name        string
		createPage  bool
		pageContent string
	}{
		{"success", true, "test"},
		{"file not found", false, ""},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			if tc.createPage {
				err := os.WriteFile(fmt.Sprintf("%s/index.html", dir), []byte(tc.pageContent), 0644)
				if err != nil {
					t.Fatalf("could not create file: %v", err)
				}
			}

			view := HTMLView{pagesDir: dir}

			data, err := view.Index()

			if tc.createPage {
				assert.NoError(t, err)
				assert.Equal(t, tc.pageContent, string(data))
			} else {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, fs.ErrNotExist))
			}
		})
	}
}

func TestError(t *testing.T) {
	view := HTMLView{}

	data, err := view.Error(context.Background(), "test")

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Contains(t, string(data), "test")
}
