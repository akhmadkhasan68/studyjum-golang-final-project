package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("file .env not found")
	}

	return os.Getenv(key)
}
