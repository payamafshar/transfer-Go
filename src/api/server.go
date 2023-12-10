package api

import (
	accountroutes "ReservApp/src/api/account/accountRoutes"
	authroutes "ReservApp/src/api/auth/authRoutes"
	transferroutes "ReservApp/src/api/transfer/transfer_routes"
	"ReservApp/src/cmd"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupServer(cfg *cmd.AppConfig) error {
	server := gin.Default()
	//set Mode of application
	gin.SetMode(gin.DebugMode)
	// binding custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", cmd.ValidCurrency)
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	setupRoutes(server)
	return server.Run(fmt.Sprintf("0.0.0.0:%d", cfg.Api.ApiPort))
}

func setupRoutes(server *gin.Engine) {
	api := server.Group("api")
	v1 := api.Group("v1")
	authroutes.SetupAuthRoutes(v1)
	accountroutes.SetupAccountRoutes(v1)
	transferroutes.SetupTransferRoutes(v1)
}
