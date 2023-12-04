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

type IAuthServices interface {
	Register(dto *dtos.RegisterUserDto) (error, *models.User)
	Login(dto *dtos.LoginDto) (string, string, string, time.Time, error)
	RefreshToken(dto *dtos.RefreshTokenDto) (string, string, string, time.Time, error)
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

type AuthService struct {
	cfg            *cmd.AppConfig
	jwtService     *jsonwebtoken.JwtService
	psqlRepository *db.PsqlRepository
}

func (s *AuthService) Register(dto *dtos.RegisterUserDto) (error, *models.User) {

	fmt.Println(dto.UserName)
	if existUser := s.psqlRepository.DB.Where("username = ?", dto.UserName).First(&models.User{}); existUser.Error == nil {

		return errors.New("user already exist"), nil
	}
	//test
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

func (s *AuthService) Login(dto *dtos.LoginDto) (string, string, string, time.Time, error) {
	var user *models.User
	if existsUser := s.psqlRepository.DB.Where("username = ?", dto.Username).First(&user); existsUser.Error != nil {
		return "", "", "", time.Time{}, errors.New("invalid credentials")
	}
	fmt.Println(*user.Password)
	isCorrectHashedPassword := helper.CompareHashWithHashString(dto.Password, *user.Password)
	if !isCorrectHashedPassword {
		return "", "", "", time.Time{}, errors.New("invalid credentials")
	}
	// passing nil to expirationdat and will get default value
	accessToken, tokenErr := s.jwtService.GenrateJwtToken(*user.Username, &s.cfg.Api.JWTtokenExpiresDate)
	if tokenErr != nil {
		return "", "", "", time.Time{}, tokenErr
	}
	refreshToken := s.jwtService.GenerateRefreshToken()
	user.RefreshToken = &refreshToken
	refreshTokenExpireDate := time.Now().AddDate(0, 0, s.cfg.Api.JWTrefreshTokenExpiresDate)
	user.RefreshTokenExpireDate = &refreshTokenExpireDate
	s.psqlRepository.DB.Save(user)
	return *accessToken, refreshToken, *user.Username, refreshTokenExpireDate, nil

}
func (s *AuthService) RefreshToken(dto *dtos.RefreshTokenDto) (string, string, string, time.Time, error) {

	accessToken := dto.Token
	refreshToken := dto.RefreshToken

	user, err := s.jwtService.VerifyToken(accessToken)
	if err != nil {
		return "", "", "", time.Time{}, err
	}

	if *user.RefreshToken != refreshToken || (*user.RefreshTokenExpireDate).Before(time.Now()) {
		return "", "", "", time.Time{}, errors.New("invalid client Request")
	}
	fmt.Println(refreshToken)
	//passing refreshToken expiresDate
	jwtRefreshToken, errR := s.jwtService.GenrateJwtToken(*user.Username, &s.cfg.Api.JWTrefreshTokenExpiresDate)
	jwtAccessToken, errA := s.jwtService.GenrateJwtToken(*user.Username, &s.cfg.Api.JWTtokenExpiresDate)
	if errA != nil || errR != nil {
		return "", "", "", time.Time{}, err
	}
	refreshTokenExpireDate := time.Now().AddDate(0, 0, s.cfg.Api.JWTrefreshTokenExpiresDate)

	user.RefreshToken = jwtRefreshToken
	user.RefreshTokenExpireDate = &refreshTokenExpireDate
	s.psqlRepository.DB.Save(user)
	return *jwtAccessToken, *jwtRefreshToken, *user.Username, refreshTokenExpireDate, nil
}
