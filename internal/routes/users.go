package routes

import (
	"go-auth/internal/handler"

	"github.com/gin-gonic/gin"
)

func addUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/", handler.CreateUser)
	users.GET("/", handler.GetUsers)
	users.GET("/:id", handler.GetUserById)
	users.PATCH("/:id", handler.UpdateUser)
	users.DELETE("/:id", handler.DeleteUser)
}
