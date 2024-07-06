package config

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"slack.app/internal/domain"
)

func InitMongoDB(cfg AppConfig) {
	err := mgm.SetDefaultConfig(nil, cfg.MongodbDatabase, options.Client().ApplyURI(cfg.MongoLink))
	if err != nil {
		panic(err)
	}
	domain.Init()
	log.Println("Connected to MongoDB!")
}
