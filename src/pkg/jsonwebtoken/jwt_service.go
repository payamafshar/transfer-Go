package jsonwebtoken

import (
	"ReservApp/src/cmd"
	"ReservApp/src/db"
	"crypto/rand"
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

func (s *JwtService) GenerateRefreshToken() string {
	b := make([]byte, 24)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
