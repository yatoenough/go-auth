package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// define config variables
var (
	port         int
	apiHost      string
	databaseURL  string
	databaseName string
	senderMail   string
	appPassword  string
	smtpHost     string
	smtpPort     string
)

func LoadConfig() {
	//load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//write vars from .env to config vars
	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Port must be a number.")
	}
	apiHost = os.Getenv("API_HOST")
	databaseURL = os.Getenv("DB_URL")
	databaseName = os.Getenv("DB_NAME")
	senderMail = os.Getenv("MAILER_SENDER_MAIL")
	appPassword = os.Getenv("MAILER_APP_PASSWORD")
	smtpHost = os.Getenv("MAILER_SMTP_HOST")
	smtpPort = os.Getenv("MAILER_SMTP_PORT")
}

func GetPort() int {
	return port
}

func GetApiHost() string {
	return apiHost
}

func GetDatabaseURL() string {
	return databaseURL
}

func GetDatabaseName() string {
	return databaseName
}

func GetSenderMail() string {
	return senderMail
}

func GetAppPassword() string {
	return appPassword
}

func GetSmtpHost() string {
	return smtpHost
}

func GetSmtpPort() string {
	return smtpPort
}
