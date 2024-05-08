package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	//load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
