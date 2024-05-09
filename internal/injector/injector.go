package injector

import (
	"context"
	"go-auth/internal/controller"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/service"
)

var (
	userService    service.UserService
	UserController controller.UserController
)

func Init() {
	userService = service.NewUserService(mongodb.UsersCollection, context.TODO())
	UserController = controller.New(userService)
}
