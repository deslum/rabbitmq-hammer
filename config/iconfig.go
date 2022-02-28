package config

type IConfig interface {
	GetConfig() (*IConfig, error)
	String() string
}
