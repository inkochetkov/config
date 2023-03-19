package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// NewConfig returns a new decoded сonfig struct by layout
func NewConfig(configPath string, entity any) (any, error) {

	config := entity

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
