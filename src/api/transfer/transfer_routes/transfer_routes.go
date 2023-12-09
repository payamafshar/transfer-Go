package transferroutes

import (
	"ReservApp/src/api/transfer"

	"github.com/gin-gonic/gin"
)

func SetupTransferRoutes(group *gin.RouterGroup) {
	transferRoute := group.Group("transfer")
	transferHandler := transfer.NewTransferHandler()

	transferRoute.POST("/create", transferHandler.Create)
}
