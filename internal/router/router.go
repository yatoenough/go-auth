package router

import (
	"fmt"
	"go-auth/internal/injector"
	"go-auth/internal/middleware"
	"go-auth/internal/model/dto"
	"go-auth/internal/router/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	v1.GET("/protected", middleware.AuthMiddleware, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": ctx.MustGet("token_data")})
	})

	routes.RegisterUserRoutes(v1)
	routes.RegisterAuthRoutes(v1)
}
