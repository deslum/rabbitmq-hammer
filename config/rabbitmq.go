package config

import "fmt"

type RabbitMQ struct {
	URI string `json:"URI"`
}

func (o RabbitMQ) String() string {
	return fmt.Sprintf(
		`
		URI: %v`,
		o.URI)
}
