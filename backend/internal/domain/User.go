package domain

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string    `json:"email" bson:"email"`
	Username         string    `json:"username" bson:"username"`
	Password         string    `json:"-" bson:"password"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

func NewUser(email string, password string, username string) *User {
	return &User{
		Email:     email,
		Password:  password,
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (model *User) CollectionName() string {
	return "users"
}

func Init() {
	ctx := context.Background()
	coll := mgm.Coll(&User{})
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "username", Value: 1},
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := coll.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(err)
	}
}
