package config

import "time"

type Server struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port" default:"8080"`
	Timeout Timeout
	Other   Other
}

type Timeout struct {
	Request time.Duration `yaml:"write" default:"10s"`
	Write   time.Duration `yaml:"write" default:"1m"`
	Read    time.Duration `yaml:"read" default:"1m"`
}
