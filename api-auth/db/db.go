package db

import (
	"github.com/api-auth/models"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("stroage.db"))

	if err != nil {
		log.Fatal("failed to connect to the DB", err)
	}

	db.AutoMigrate(&models.User{}, &models.Book{})
	return db
}
