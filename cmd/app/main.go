package main

import (
	"go-auth/config"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/injector"
	"go-auth/internal/router"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.LoadConfig()

	err := mongodb.Init(os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	injector.InitDependencies()

	defer func() {
		err := mongodb.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router.Run(os.Getenv("PORT"))
}
