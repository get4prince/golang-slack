package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (AppConfig, error) {
	fmt.Printf("%v\n", os.Getenv("APP_ENV"))
	godotenv.Load()

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{
		ServerPort: httpPort,
	}, nil

}
