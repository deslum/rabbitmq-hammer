package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"rabbitmq-hammer/compression"
)

// RMQHammerConfig app configuration struct
type RMQHammerConfig struct {
	RabbitMQ `json:"RabbitMQ"`

	Publisher    *Publisher               `json:"Producer"`
	Consumer     *Consumer                `json:"Consumer"`
	CompressType compression.CompressType `json:"CompressType"`
}

// NewRMQHammerConfig init struct RMQHammerConfig.
func NewRMQHammerConfig() *RMQHammerConfig {
	return &RMQHammerConfig{}
}

// GetConfig read and get app configuration.
func (o *RMQHammerConfig) GetConfig() (*RMQHammerConfig, error) {
	// ReadConfig read config file RMQHammer
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile(%v) error %v", "config.json", err)
	}

	err = json.Unmarshal(data, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// TODO Config input in console
func (o RMQHammerConfig) String() string {
	return fmt.Sprintf(`
RabbitMQ-Hammer config: 
	RabbitMQ: %v`, o.RabbitMQ)
}
