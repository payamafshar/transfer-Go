package auth

import (
	"ReservApp/src/cmd"
)

type AuthService struct {
	cfg *cmd.AppConfig
	// userRepo      *repositories.UserRepository
	// JwtService    *JwtService
}

func NewAuthService(cfg *cmd.AppConfig) *AuthService {

	return &AuthService{
		cfg: cfg,
	}
}

// func (service *AuthService) Register(dto *dtos.RegisterUserDto) error {

// }
