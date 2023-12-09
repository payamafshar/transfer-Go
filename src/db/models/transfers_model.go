package models

import (
	"time"

	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	Id            int        `json:"id" gorm:"primaryKey autoIncrement"`
	FromAccountID string     `json:"from_accountId"`
	FromAccount   Account    `gorm:"foreignKey:FromAccountID"`
	ToAccountID   string     `json:"to_accountId"`
	ToAccount     Account    `gorm:"foreignKey:ToAccountID"`
	Amount        int64      `json:"amount"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	Entrie        []Entrie   `gorm:"foreignKey:TransferID"`
}

//belongs_to is used from the owner's side (the user in this case), while has_one is used from the owned side (the profile in this case)
