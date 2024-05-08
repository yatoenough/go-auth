package main

import (
	"go-auth/config"
	"go-auth/internal/database"
	"go-auth/internal/routes"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.LoadConfig()

	err := database.Init(config.DatabaseURL, config.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := database.CLose()
		if err != nil {
			log.Fatal(err)
		}
	}()

	routes.Run(config.Port)
}
