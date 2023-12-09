package dtos

type CreateAccountDto struct {
	Currency string `json:"currency,omitempty" binding:"required,oneof=EUR USD"`
}
type UpdateAccountDto struct {
	Currency string `json:"currency,omitempty" binding:"oneof=EUR USD"`
}
