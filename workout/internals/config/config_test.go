package config

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitThrowsErrorWhenNoEnvFile(t *testing.T) {
	tempDir := t.TempDir()
	os.Chdir(tempDir)

	defer func(t *testing.T) {
		if r := recover(); r != nil {
			assert.NotNil(t, r)
		}
	}(t)

	Init()

	assert.FailNow(t, "Init() did not panic")
}

func TestInitWorksWhenEnvFile(t *testing.T) {
	tempDir := t.TempDir()
	os.Chdir(tempDir)
	data := []byte(`TEST=val`)
	os.WriteFile(".env", data, fs.FileMode(0777))

	defer func() {
		if r := recover(); r != nil {
			switch ty := r.(type) {
			case string:
				assert.FailNow(t, ty)
			case error:
				assert.FailNow(t, ty.Error())
			default:
				assert.FailNow(t, "Init() panicked")
			}
		}
	}()

	Init()
}

func TestEnv(t *testing.T) {
	tempDir := t.TempDir()
	os.Chdir(tempDir)
	data := []byte(`TEST=val`)
	os.WriteFile(".env", data, fs.FileMode(0777))
	Init()

	testcases := []struct {
		name     string
		key      string
		expected string
	}{
		{"test value not found", "no-val", ""},
		{"test value found", "TEST", "val"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			val := Env(tc.key)

			assert.Equal(t, tc.expected, val)
		})
	}
}
