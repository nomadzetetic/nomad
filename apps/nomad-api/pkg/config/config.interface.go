package config

type GinContextKey string

type Service interface {
	GetPort() string
	GetGinContextKey() GinContextKey
	GetPostgresDatabaseUrl() string
	GetJwtSecret() string
}
