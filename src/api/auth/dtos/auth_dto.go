package dtos

type RegisterUserDto struct {
	UserName string `json:"user_name,omitempty" validate:"min=5,max=20,required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=16,required"`
}
