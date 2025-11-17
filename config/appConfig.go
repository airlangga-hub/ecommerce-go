package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)


type AppConfig struct {
	ServerPort string
}


func SetupEnv() (AppConfig, error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		return AppConfig{}, errors.New("http port env variable not found")
	}

	return AppConfig{ServerPort: httpPort}, nil
}