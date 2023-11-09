package auth

import "ReservApp/src/db/models"

type AuthenticateReponse struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterResponse struct {
	User *models.User
}
