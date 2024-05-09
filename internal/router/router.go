package router

import (
	"fmt"
	"go-auth/internal/injector"
	"go-auth/internal/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// init router
var router = gin.Default()

func Run(port string) {
	registerRoutes()
	router.NoRoute(func(c *gin.Context) {
		dto.ApiResponse(c, http.StatusNotFound, "Resource not found.")
	})
	router.Run(fmt.Sprintf(":%s", port))
}

func registerRoutes() {
	router.GET("/activate/:code", injector.UserController.ActivateUser)

	v1 := router.Group("/api/v1")
	injector.UserController.RegisterUserRoutes(v1)
}
