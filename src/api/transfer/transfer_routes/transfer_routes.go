package transferroutes

import (
	"ReservApp/src/api/transfer"
	"ReservApp/src/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupTransferRoutes(group *gin.RouterGroup) {
	transferRoute := group.Group("transfer")
	transferHandler := transfer.NewTransferHandler()

	transferRoute.POST("/create", middlewares.AuthorizationMiddleware(), transferHandler.Create)
}
