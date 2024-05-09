package routes

import (
	"go-auth/internal/injector"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	authRouter := rg.Group("/auth")

	authRouter.POST("/register", injector.AuthController.Register)
	authRouter.POST("/login", injector.AuthController.Login)
}
