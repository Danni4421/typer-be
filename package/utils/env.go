package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	environmentError := godotenv.Load()

	if environmentError != nil {
		panic("Error loading .env file")
	}
}

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
