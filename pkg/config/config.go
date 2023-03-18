package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// root config
type Config struct {
	HTTP     *Server   `yaml:"server_http"`
	GRPC     *Server   `yaml:"server_grpc"`
	DataBase *DataBase `yaml:"data_base"`
	Broker   *Broker   `yaml:"broker"`
	Other    Other
}

type Other map[string]any

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {

	config := &Config{}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
