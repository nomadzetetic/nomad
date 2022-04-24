package config

import (
	"log"
	"os"
)

type EnvConfigService struct{}

func (config EnvConfigService) GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return port
}

func (config EnvConfigService) GetGinContextKey() GinContextKey {
	return "GinContext"
}

func (config EnvConfigService) GetPostgresDatabaseUrl() string {
	dbUrl := os.Getenv("POSTGRES_DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("POSTGRES_DATABASE_URL env variable not set")
	}
	return dbUrl
}

func (config EnvConfigService) GetJwtSecret() string {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET env variable not set")
	}
	return jwtSecret
}
