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

	err := mongodb.Init(config.DatabaseURL, config.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := mongodb.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	routes.Run(config.Port)
}
