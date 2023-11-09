package dtos

type RegisterUserDto struct {
	UserName string `json:"user_name,omitempty" validate:"min=5,max=20,required"`
	Password string `json:"password" validate:"min=4,max=16,required"`
}

type LoginDto struct {
	Username string `json:"user_name,omitempty" validate:"min=5,max=20,required"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenDto struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
