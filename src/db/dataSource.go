package db

import (
	"ReservApp/src/db/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PsqlRepository struct {
	DB *gorm.DB
}

func NewPsqlRepository() *PsqlRepository {
	db := Init()
	return &PsqlRepository{db}
}

func Init() *gorm.DB {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Book{})
	return db
}
