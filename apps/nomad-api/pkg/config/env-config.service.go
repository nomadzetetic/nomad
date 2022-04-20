package config

import "os"

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
