package config

type GinContextKey string

type ConfigService interface {
	GetPort() string
	GetGinContextKey() GinContextKey
}
