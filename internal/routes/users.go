package routes

import (
	"go-auth/internal/handler/users"

	"github.com/gin-gonic/gin"
)

func addUsersRoutes(rg *gin.RouterGroup) {
	//define route group
	usersGroup := rg.Group("/users")

	//bind methods
	usersGroup.POST("/", users.CreateUser)
	usersGroup.GET("/", users.GetUsers)
	usersGroup.GET("/:id", users.GetUserById)
	usersGroup.PATCH("/:id", users.UpdateUser)
	usersGroup.DELETE("/:id", users.DeleteUser)
}
