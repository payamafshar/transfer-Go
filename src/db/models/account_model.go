package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Id        int        `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Owner     string     `json:"owner"`
	Balance   int64      `json:"balance"`
	Currency  string     `json:"currency"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
}
