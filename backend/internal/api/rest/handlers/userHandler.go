package handlers

import (
	"github.com/gin-gonic/gin"
	"slack.app/config"
	"slack.app/internal/controllers"
)

func ApplyMiddleware(cfg config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cfg)
		c.Next()
	}
}

func SetupUserRoutes(router *gin.Engine, cfg config.AppConfig) {
	router.Use(ApplyMiddleware(cfg))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	router.POST("/register", controllers.RegisterHandlers)
	router.POST("/login", controllers.LoginHandlers)
}
