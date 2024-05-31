package config

import (
	"fmt"
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
	key := "TEST"
	expected := "val"
	data := []byte(fmt.Sprintf("%s=%s", key, expected))
	os.WriteFile(".env", data, fs.FileMode(0777))
	Init()

	val := Env(key)

	assert.Equal(t, expected, val)
}

func TestEnvPanicsIfValueNotFound(t *testing.T) {
	tempDir := t.TempDir()
	os.Chdir(tempDir)
	data := []byte("TEST=val")
	os.WriteFile(".env", data, fs.FileMode(0777))
	Init()

	key := "panics"
	assert.PanicsWithError(t, ErrNotFound(key).Error(), func() {
		Env(key)
	})
}
