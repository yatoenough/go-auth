package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run(port int) {
	registerRoutes()
	router.Run(fmt.Sprintf(":%d", port))
}

func registerRoutes() {
	v1 := router.Group("/api/v1")
	// authenticated := v1.Group("/authenticated")
	addUsersRoutes(v1)
}
