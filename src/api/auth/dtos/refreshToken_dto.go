package dtos

type RefreshTokenDto struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
