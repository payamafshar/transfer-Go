package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}
