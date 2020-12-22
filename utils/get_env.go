package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvWithKey => to get value from env
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

// LoadEnv => to load .env
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
}
