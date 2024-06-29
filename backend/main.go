package main

import (
	"fmt"
	"log"

	"slack.app/config"
)

func main() {

	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("Config File Not Loaded %v\n", err)
	}
	fmt.Printf("AppConfig: %+v\n", cfg)
}
