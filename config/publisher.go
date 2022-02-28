package config

type Publisher struct {
	Message       Message `json:"Message"`
	TTL           uint64  `json:"TTL"`
	HammersCount  int     `json:"HammersCount"`
	MessagesCount uint64  `json:"MessagesCount"`
	Exchange      string  `json:"Exchange"`
	RoutingKey    string  `json:"RoutingKey"`
}
