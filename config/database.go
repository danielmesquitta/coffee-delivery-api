package config

import (
	"log"
	"os"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectToDatabase() {
	var err error

	dsn := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
}

func autoMigrate() {
	db.AutoMigrate(&model.Product{})
}

func GetDatabase() *gorm.DB {
	if db == nil {
		log.Fatal("database not initialized")
	}

	return db
}
