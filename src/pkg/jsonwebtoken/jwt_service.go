package jsonwebtoken

import (
	"ReservApp/src/cmd"
	"ReservApp/src/db"
	"ReservApp/src/db/models"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	cfg            *cmd.AppConfig
	psqlRepository *db.PsqlRepository
}

func NewJwtService(cfg *cmd.AppConfig, r *db.PsqlRepository) *JwtService {
	return &JwtService{
		cfg:            cfg,
		psqlRepository: r,
	}
}

func (s *JwtService) GenrateJwtToken(identifire string, exp time.Duration) (*string, error) {

	claims := &AccessTokenClaims{
		Identifier: identifire,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.Api.JWTSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, err
}

func (s *JwtService) VerifyToken(tokenString string) (*models.User, error) {
	var user *models.User
	// Extract and verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.Api.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if token is valid
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["exp"].(float64) < float64(time.Now().Unix()) {
		return nil, errors.New("invalid or expired token")
	}

	// Check if user exists
	clm := claims["identifier"].(string)

	if existsUser := s.psqlRepository.DB.Where("username = ?", clm).First(&user); existsUser.Error != nil {
		return nil, errors.New("invalid or expired token")
	}

	return user, nil
}

func (s *JwtService) GenerateRefreshToken() string {
	b := make([]byte, 24)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
