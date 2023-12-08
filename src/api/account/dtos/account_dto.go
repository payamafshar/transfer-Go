package dtos

type CreateAccountDto struct {
	Currency string `json:"currency,omitempty" binding:"required,oneof=EUR USD"`
}
