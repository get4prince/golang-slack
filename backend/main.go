package main

import (
	"log"

	"slack.app/config"
	api "slack.app/internal/api/rest"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("Config File Not Loaded %v\n", err)
	}
	config.InitMongoDB(cfg)
	api.StartServer(cfg)

}
