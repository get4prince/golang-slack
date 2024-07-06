package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
	"slack.app/config"
	"slack.app/internal/services"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (li LoginInput) ValidateLoginInput() error {
	return validation.ValidateStruct(&li,
		validation.Field(&li.Username, validation.Required),
		validation.Field(&li.Password, validation.Required),
	)
}

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (ri RegisterInput) ValidateRegisterInput() error {
	return validation.ValidateStruct(&ri,
		validation.Field(&ri.Username, validation.Required),
		validation.Field(&ri.Password, validation.Required),
		validation.Field(&ri.Email, validation.Required, is.EmailFormat),
	)
}

func LoginHandlers(c *gin.Context) {
	var requestBody LoginInput
	config, ok := c.MustGet("config").(config.AppConfig)
	if !ok {
		c.JSON(400, gin.H{"error": ok})
		return
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	validateData := requestBody.ValidateLoginInput()
	if validationErrs := requestBody.ValidateLoginInput(); validationErrs != nil {
		c.JSON(400, gin.H{"error": validationErrs})
		return
	}

	user, err := services.GetUser(requestBody.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return
	}
	b, _ := json.Marshal(validateData)
	if string(b) != "null" {
		c.JSON(400, gin.H{"error": string(b)})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := services.CreateToken(user.Username, config.Jwt_key)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"token":   token,
	})
}

func RegisterHandlers(c *gin.Context) {
	var requestBody RegisterInput
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	validateData := requestBody.ValidateRegisterInput()
	b, _ := json.Marshal(validateData)
	if string(b) != "null" {
		c.JSON(400, gin.H{"error": string(b)})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := services.Register(requestBody.Email, string(password), requestBody.Username)
	fmt.Print(user, err, "user here")

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
