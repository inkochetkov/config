package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// NewConfigAny returns a new decoded the universal struct
func NewConfigAny(configPath string) (map[string]any, error) {

	config := make(map[string]any)

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
