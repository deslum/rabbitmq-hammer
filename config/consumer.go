package config

type Consumer struct {
	Queues                []string `json:"Queues"`
	PrefetchCount         int      `json:"PrefetchCount"`
	AutoAck               bool     `json:"AutoAck"`
	Exclusive             bool     `json:"Exclusive"`
	NoLocal               bool     `json:"NoLocal"`
	NoWait                bool     `json:"NoWait"`
	LatencyInMicroseconds uint64   `json:"LatencyInMicroseconds"`
	ConsumersCount        int      `json:"ConsumersCount"`
}
