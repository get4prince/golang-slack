package services

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"slack.app/internal/domain"
)

func Register(email string, password string, username string) (*domain.User, error) {
	user := domain.NewUser(email, password, username)
	err := mgm.Coll(user).Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(username string) (*domain.User, error) {
	user := &domain.User{}
	filter := bson.D{{"username", username}}
	err := mgm.Coll(user).FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return user, nil
}

func CreateToken(username string, jwtKey string) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: expirationTime.Unix(),
	}

	// Create the token using the claims and the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
