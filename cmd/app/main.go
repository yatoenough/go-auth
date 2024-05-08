package main

import (
	"cruddemo/internal/database"
	"cruddemo/internal/routes"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
    connectionString := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	err := database.Init(connectionString, dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := database.CLose()
		if err != nil {
			log.Fatal(err)
		}
	}()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	routes.Run(port)
}
