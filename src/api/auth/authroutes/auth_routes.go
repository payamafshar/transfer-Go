package authroutes

import (
	auth "ReservApp/src/api/auth"
	"ReservApp/src/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(group *gin.RouterGroup) {
	authRoute := group.Group("auth")
	authHandler := auth.NewAuthHandler()

	authRoute.POST("/register", authHandler.Register)

	authRoute.POST("/login", authHandler.Login)

	authRoute.POST("/refreshToken", authHandler.RefreshToken)
	authRoute.GET("/current", middlewares.AuthorizationMiddleware(), authHandler.CurrentUser)

}
