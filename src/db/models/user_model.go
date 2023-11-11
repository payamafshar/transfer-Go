package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id                     *int       `json:"id" gorm:"primaryKey autoIncrement"`
	Username               *string    `gorm:"type:varchar(40);unique" json:"username"`
	Password               *string    `json:"passowrd"`
	RefreshToken           *string    `json:"refresh_token"`
	RefreshTokenExpireDate *time.Time `json:"refreshToken_expiresDate"`
}
