package main

import (
	"cruddemo/internal/database"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := database.Init(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer func(){
		err := database.CLose()
		if err != nil {
			log.Fatal(err)
		}
	}()

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	r := gin.Default()
	r.Run(fmt.Sprintf(":%d", port))
}
