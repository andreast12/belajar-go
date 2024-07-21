package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func IsInProduction() bool {
	ginMode := os.Getenv("GIN_MODE")
	return ginMode == "release"
}