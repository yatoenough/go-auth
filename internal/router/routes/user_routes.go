package routes

import (
	"go-auth/internal/injector"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoutes := rg.Group("/users")

	userRoutes.GET("/", injector.UserController.GetAll)
	userRoutes.GET("/:id", injector.UserController.GetUserById)
	userRoutes.DELETE("/:id", injector.UserController.DeleteUser)
}
