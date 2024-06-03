package router

import (
	"fmt"
	"go-auth/internal/injector"
	"go-auth/internal/model/dto"
	"go-auth/internal/router/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run(port string) {
	registerRoutes()

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}

func registerRoutes() {
	router.NoRoute(func(c *gin.Context) {
		dto.ApiResponse(c, http.StatusNotFound, "Resource not found.")
	})
	router.GET("/activate/:code", injector.UserController.ActivateUser)

	v1 := router.Group("/api/v1")
	v1.GET("/protected", injector.AuthMiddleware.Guard, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": ctx.MustGet("token_data")})
	})

	routes.RegisterUserRoutes(v1)
	routes.RegisterAuthRoutes(v1)
}
