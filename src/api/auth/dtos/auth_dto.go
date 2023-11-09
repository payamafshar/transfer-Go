package dtos

type RegisterUserDto struct {
	UserName string `json:"user_name,omitempty" validate:"min=5,max=20,required"`
	Password string `json:"password" validate:"min=4,max=16,required"`
}

type LoginDto struct {
	Username string `json:"user_name,omitempty" validate:"min=5,max=20,required"`
	Password string `json:"password" validate:"required"`
}
