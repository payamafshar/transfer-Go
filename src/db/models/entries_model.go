package models

import (
	"time"

	"gorm.io/gorm"
)

type Entrie struct {
	gorm.Model
	Id        int        `json:"id" gorm:"primaryKey autoIncrement"`
	AccountId string     `json:"account_id"`
	Amount    int64      `json:"amount"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
}
