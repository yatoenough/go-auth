package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port         int
	DatabaseURL  string
	DatabaseName string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Port must be a number.")
	}
	DatabaseURL = os.Getenv("DB_URL")
	DatabaseName = os.Getenv("DB_URL")
}
