package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort      string
	MongoLink       string
	MongodbDatabase string
	Jwt_key         string
}

func SetupEnv() (AppConfig, error) {
	fmt.Printf("%v\n", os.Getenv("APP_ENV"))
	godotenv.Load()
	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	mongoLink := os.Getenv("MONGO_DB")
	if len(mongoLink) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	MongodbDatabase := os.Getenv("MONGO_DATABASE")
	if len(MongodbDatabase) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	JwtKey := os.Getenv("JWT_KEY")
	if len(MongodbDatabase) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{
		ServerPort:      httpPort,
		MongoLink:       mongoLink,
		MongodbDatabase: MongodbDatabase,
		Jwt_key:         JwtKey,
	}, nil
}

func InitServer() *gin.Engine {
	r := gin.Default()
	r.Run(":4000")
	return r
}
