package test

import (
	"testing"

	"github.com/inkochetkov/config/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Log("Test check config")

	conf, err := config.NewConfig("config.yaml")
	assert.NoError(t, err)

	want := &config.Config{
		HTTP: &config.Server{
			Host: "127.0.0.1",
			Port: "8080",
		},
	}

	have := conf
	assert.Equal(t, want, have)

	t.Log("Test saccess")
}
