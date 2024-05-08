package main

import (
	"go-auth/config"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/routes"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.LoadConfig()

	//init db
	err := mongodb.Init(os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	//close connection to db after app stops
	defer func() {
		err := mongodb.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//start app
	routes.Run(os.Getenv("PORT"))
}
