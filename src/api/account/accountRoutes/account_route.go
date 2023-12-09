package accountroutes

import (
	"ReservApp/src/api/account"
	"ReservApp/src/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupAccountRoutes(group *gin.RouterGroup) {
	accountRoute := group.Group("account")
	accountHandler := account.NewAccountHandler()

	accountRoute.POST("/create", middlewares.AuthorizationMiddleware(), accountHandler.Create)
	accountRoute.GET("/test", middlewares.AuthorizationMiddleware(), accountHandler.FindAll)
	accountRoute.GET("/:id", accountHandler.GetById)
	accountRoute.PUT("/:id", middlewares.AuthorizationMiddleware(), accountHandler.Update)
	accountRoute.DELETE("/:id", middlewares.AuthorizationMiddleware(), accountHandler.Delete)
}
