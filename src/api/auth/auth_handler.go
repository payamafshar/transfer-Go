package auth

import (
	"ReservApp/src/api/auth/dtos"
	"ReservApp/src/cmd"
	"fmt"
	"net/http"

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

func (h *AuthHandler) Register(ctx *gin.Context) {
	dto := new(dtos.RegisterUserDto)
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//calling service
	creationError, user := h.authService.Register(dto)
	fmt.Println(user)
	if creationError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": creationError.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, &RegisterResponse{User: user})
	return
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	dto := new(dtos.LoginDto)
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, refreshToken, userName, refreshTokenExpireDate, loginErr := h.authService.Login(dto)

	if loginErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": loginErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &AuthenticateReponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: userName, RefreshTokenExpiresDate: refreshTokenExpireDate})
	return
}

func (h *AuthHandler) CurrentUser(ctx *gin.Context) {
	user := ctx.MustGet("claims")
	ctx.JSON(http.StatusOK, user)
	return
}
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {

	dto := new(dtos.RefreshTokenDto)
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, refreshToken, userName, refreshTokenExpireDate, loginErr := h.authService.RefreshToken(dto)
	if loginErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": loginErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &AuthenticateReponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: userName, RefreshTokenExpiresDate: refreshTokenExpireDate})
	return
}
