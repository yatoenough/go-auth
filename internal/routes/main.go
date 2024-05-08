package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Run will start the server
func Run(port int) {
	getRoutes()
	router.Run(fmt.Sprintf(":%d", port))
}

func getRoutes() {
	v1 := router.Group("/api/v1")
	addUsersRoutes(v1)
}
