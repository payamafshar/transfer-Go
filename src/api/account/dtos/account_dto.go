package dtos

type CreateAccountDto struct {
	Currency string `json:"currency,omitempty" validate:"required_with=USD EUR"`
}
