package config

import "time"

type DataBase struct {
	URL     string
	Timeout struct {
		MaxOpenConns    int           `yaml:"max_open_conns" default:"0"`
		MaxIdleConns    int           `yaml:"max_idle_conns" default:"2"`
		ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" default:"1m"`
	}
	Other Other
}
