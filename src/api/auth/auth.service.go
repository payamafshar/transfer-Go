package auth

import (
	"ReservApp/src/cmd"
	"ReservApp/src/db"
)

type AuthService struct {
	cfg *cmd.AppConfig
	// userRepo      *repositories.UserRepository
	// JwtService    *JwtService
	psqlRepository *db.PsqlRepository
}

func NewAuthService(cfg *cmd.AppConfig) *AuthService {
	psqlRepository := db.NewPsqlRepository()
	return &AuthService{
		cfg:            cfg,
		psqlRepository: psqlRepository,
	}
}

// func (service *AuthService) Register(dto *dtos.RegisterUserDto) error {

// }
