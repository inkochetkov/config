package test

import (
	"testing"

	"github.com/inkochetkov/config/pkg/config"
	"github.com/stretchr/testify/assert"
)

type Config struct {
	HTTP *Server `yaml:"server_http"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port" default:"8080"`
}

func TestConfig(t *testing.T) {
	t.Log("Test check config")

	conf, err := config.NewConfig("config.yaml", &Config{})
	assert.NoError(t, err)

	want := &Config{
		HTTP: &Server{
			Host: "127.0.0.1",
			Port: "8080",
		},
	}

	have := conf
	assert.Equal(t, want, have)

	t.Log("Test saccess")
}
