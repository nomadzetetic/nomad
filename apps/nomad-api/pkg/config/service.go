package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type GinContextKey string

type Service interface {
	GetPort() string
	GetGinContextKey() GinContextKey
	GetPostgresDatabaseUrl() string
	GetJwtSecret() string
}

type EnvConfig struct{}

func NewConfig() *EnvConfig {
	_ = godotenv.Load(".env", "../.env")
	return &EnvConfig{}
}

func (c EnvConfig) GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return port
}

func (c EnvConfig) GetGinContextKey() GinContextKey {
	return "GinContext"
}

func (c EnvConfig) GetPostgresDatabaseUrl() string {
	dbUrl := os.Getenv("POSTGRES_DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("POSTGRES_DATABASE_URL env variable not set")
	}
	return dbUrl
}

func (c EnvConfig) GetJwtSecret() string {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET env variable not set")
	}
	return jwtSecret
}
