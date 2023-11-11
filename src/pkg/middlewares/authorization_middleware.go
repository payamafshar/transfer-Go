package middlewares

import (
	"ReservApp/src/cmd"
	"ReservApp/src/db"
	"ReservApp/src/pkg/jsonwebtoken"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var once sync.Once
var cfg *cmd.AppConfig
var psqlRepository *db.PsqlRepository
var jwtService *jsonwebtoken.JwtService

func AuthorizationMiddleware() gin.HandlerFunc {

	once.Do(func() {
		psqlRepository = db.NewPsqlRepository()
		cfg = cmd.GetAppConfig()
		jwtService = jsonwebtoken.NewJwtService(cfg, psqlRepository)
	})
	return func(ctx *gin.Context) {
		headerToken := ctx.GetHeader("Authorization")
		var token string
		if headerToken == "" {
			unauthorized(ctx)
			return
		}
		splitedToken := strings.Split(headerToken, "Bearer ")
		token = splitedToken[1]

		user, err := jwtService.VerifyToken(token)

		if err != nil {
			unauthorized(ctx)
			return
		}
		ctx.Set("claims", user)
		ctx.Next()
	}
}

func unauthorized(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": "Unauthorized",
	})
}
