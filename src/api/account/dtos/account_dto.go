package dtos

type CreateAccountDto struct {
	Currency string `json:"currency,omitempty" binding:"required,currency"`
}
type UpdateAccountDto struct {
	Currency string `json:"currency,omitempty" binding:"currency"`
}
