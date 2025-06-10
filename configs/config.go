package configs

import (
	"log"
	"rogeriods/fiber-jwt-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	JWTSecret = []byte("2baad95bafc04cc1a8e4d1e292a782147a74d5dbaa1ef59bc2e533fda7c278ab")
)

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("mydata.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrations
	DB.AutoMigrate(&models.User{})
}
