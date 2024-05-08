package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// init router
var router = gin.Default()

func Run(port string) {
	registerRoutes()
	router.Run(fmt.Sprintf(":%s", port))
}

func registerRoutes() {
	//define root route group
	v1 := router.Group("/api/v1")
	addUsersRoutes(v1)
}
