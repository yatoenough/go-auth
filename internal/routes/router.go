package routes

import (
	"fmt"
	"go-auth/internal/handler"

	"github.com/gin-gonic/gin"
)

// init router
var router = gin.Default()

func Run(port int) {
	registerRoutes()
	router.Run(fmt.Sprintf(":%d", port))
}

func registerRoutes() {
	//define route for account activation
	router.GET("/activate/:link", handler.Activate)

	//define root route group
	v1 := router.Group("/api/v1")
	addUsersRoutes(v1)
}
