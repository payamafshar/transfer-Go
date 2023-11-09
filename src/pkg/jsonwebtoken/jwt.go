package jsonwebtoken

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	Identifier string `json:"identifier"`
	jwt.RegisteredClaims
}
