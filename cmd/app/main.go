package main

import (
	"go-auth/config"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/routes"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.LoadConfig()

	//init db
	err := mongodb.Init(config.GetDatabaseURL(), config.GetDatabaseName())
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
	routes.Run(config.GetPort())
}
