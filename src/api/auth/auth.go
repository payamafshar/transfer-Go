package auth

import (
	"ReservApp/src/db/models"
	"time"
)

type AuthenticateReponse struct {
	Username                string    `json:"username"`
	AccessToken             string    `json:"access_token"`
	RefreshToken            string    `json:"refresh_token"`
	RefreshTokenExpiresDate time.Time `json:"refreshtoken_expires_date"`
}

type RegisterResponse struct {
	User *models.User
}
