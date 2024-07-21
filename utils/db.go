package utils

import (
	"example/belajar-go/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_CONNECTION_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	// Migrate the schema
	DB.AutoMigrate(&models.User{}, &models.Todo{})
}