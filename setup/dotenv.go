package setup

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironment() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
