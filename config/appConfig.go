package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)


type AppConfig struct {
	ServerPort string
	Dsn string
}


func SetupEnv() (AppConfig, error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		return AppConfig{}, errors.New("http port env variable not found")
	}

	dsn := os.Getenv("DSN")

	if dsn == "" {
		return AppConfig{}, errors.New("dsn env variable not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: dsn}, nil
}