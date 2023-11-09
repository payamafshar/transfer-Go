package authroutes

import (
	auth "ReservApp/src/api/auth"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(group *gin.RouterGroup) {
	authRoute := group.Group("auth")
	authHandler := auth.NewAuthHandler()

	authRoute.POST("/register", authHandler.Register)

	authRoute.POST("/login", authHandler.Login)
}
