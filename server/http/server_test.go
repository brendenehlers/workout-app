package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	cfg := &ServerConfig{
		Addr: ":8080",
	}

	s := New(nil, nil, *cfg)

	assert.Equal(t, cfg.Addr, s.Addr)
}
