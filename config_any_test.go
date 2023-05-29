package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigAny(t *testing.T) {
	t.Log("Test check config any")

	conf, err := NewConfigAny("config.yaml")
	assert.NoError(t, err)

	want := map[string]any{
		"server_http": map[string]any{
			"host": "127.0.0.1",
			"port": 8080,
		},
	}

	have := conf
	assert.Equal(t, want, have)

	t.Log("Test saccess")
}
