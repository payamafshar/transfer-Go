package auth

import (
	"ReservApp/src/api/auth/dtos"
	"ReservApp/src/cmd"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	cfg         *cmd.AppConfig
	authService *AuthService
}

func NewAuthHandler() *AuthHandler {
	cfg := cmd.GetAppConfig()
	authService := NewAuthService(cfg)
	return &AuthHandler{
		cfg:         cfg,
		authService: authService,
	}
}

func (handler *AuthHandler) Register(ctx *gin.Context) {
	dto := new(dtos.RegisterUserDto)
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//calling service
	// handler.authService.Register(dto)

}
