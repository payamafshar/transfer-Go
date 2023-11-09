package api

import (
	authroutes "ReservApp/src/api/auth/authRoutes"
	"ReservApp/src/cmd"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupServer(cfg *cmd.AppConfig) error {
	server := gin.Default()
	gin.SetMode(gin.DebugMode)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	setupRoutes(server)
	return server.Run(fmt.Sprintf("localhost:%d", cfg.Api.ApiPort))
}

func setupRoutes(server *gin.Engine) {
	api := server.Group("api")
	v1 := api.Group("v1")
	authroutes.SetupAuthRoutes(v1)
}
