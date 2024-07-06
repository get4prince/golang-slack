package api

import (
	"github.com/gin-gonic/gin"
	"slack.app/config"
	"slack.app/internal/api/rest/handlers"
)

func StartServer(cfg config.AppConfig) {
	router := gin.Default()
	handlers.SetupUserRoutes(router, cfg)
	router.Run(":4000")
}
