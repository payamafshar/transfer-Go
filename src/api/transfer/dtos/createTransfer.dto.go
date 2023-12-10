package dtos

type CreateTransferDto struct {
	FromAccountId int `json:"from_accountId,omitempty" binding:"required,min=1"`
	ToAccountId   int `json:"to_accountId,omitempty" binding:"required,min=1"`
	Amount        int64
}
