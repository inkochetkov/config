package config

import "time"

type Broker struct {
	Rabbit *Rabbit
	Kafka  *Kafka
	Redis  *Redis
	MQTT   *MQTT
}

type Rabbit struct {
	Conn     ConnBroker
	Producer Producer
	Consumer Consumer
	Other    Other
}

type Kafka struct {
	Conn     ConnBroker
	Producer Producer
	Consumer Consumer
	Other    Other
}

type Redis struct {
	Conn     ConnBroker
	Producer Producer
	Consumer Consumer
	Other    Other
}

type MQTT struct {
	Conn     ConnBroker
	Producer Producer
	Consumer Consumer
	Other    Other
}

type Producer struct {
	NameTopic string
	Other     Other
}

type Consumer struct {
	NameTopic string
	Other     Other
}

type ConnBroker struct {
	URL     string        `yaml:"url"`
	Timeout time.Duration `yaml:"timeout" default:"10s"`
	Other   Other
}
