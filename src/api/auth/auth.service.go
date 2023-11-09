package auth

import (
	"ReservApp/src/api/auth/dtos"
	"ReservApp/src/cmd"
	"ReservApp/src/db"
	"ReservApp/src/db/models"
	"ReservApp/src/pkg/helper"
	"ReservApp/src/pkg/jsonwebtoken"
	"errors"
	"fmt"
	"time"
)

type AuthService struct {
	cfg            *cmd.AppConfig
	jwtService     *jsonwebtoken.JwtService
	psqlRepository *db.PsqlRepository
}

func NewAuthService(cfg *cmd.AppConfig) *AuthService {
	psqlRepository := db.NewPsqlRepository()
	jwtService := jsonwebtoken.NewJwtService(cfg, psqlRepository)
	return &AuthService{
		cfg:            cfg,
		psqlRepository: psqlRepository,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Register(dto *dtos.RegisterUserDto) (error, *models.User) {

	fmt.Println(dto.UserName)
	if existUser := s.psqlRepository.DB.Where("username = ?", dto.UserName).First(&models.User{}); existUser.Error == nil {

		return errors.New("user already exist"), nil
	}

	hashedPassword, err := helper.GenerateHash(dto.Password)
	if err != nil {
		return err, nil
	}
	user := models.User{Username: &dto.UserName, Password: &hashedPassword}
	createdUser := s.psqlRepository.DB.Model(&models.User{}).Create(&user)

	if createdUser.RowsAffected == 0 {
		return errors.New("register unsuccessfull!"), nil
	}
	return nil, &user
}

func (s *AuthService) Login(dto *dtos.LoginDto) (string, string, string, error) {
	var user *models.User
	if existsUser := s.psqlRepository.DB.Where("username = ?", dto.Username).First(&user); existsUser.Error != nil {
		return "", "", "", errors.New("invalid credentials")
	}
	fmt.Println(*user.Password)
	isCorrectHashedPassword := helper.CompareHashWithHashString(dto.Password, *user.Password)
	if !isCorrectHashedPassword {
		return "", "", "", errors.New("invalid credentials")
	}
	accessToken, tokenErr := s.jwtService.GenrateJwtToken(*user.Username, time.Duration(time.Hour*24))
	if tokenErr != nil {
		return "", "", "", tokenErr
	}
	refreshToken := s.jwtService.GenerateRefreshToken()
	user.RefreshToken = &refreshToken
	s.psqlRepository.DB.Save(user)
	return *accessToken, refreshToken, *user.Username, nil

}
