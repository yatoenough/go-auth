package injector

import (
	"context"
	"go-auth/internal/controller"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/service"
	"os"
)

var (
	userService    service.UserService
	TokenService   service.TokenService
	UserController controller.UserController
	AuthController controller.AuthController
)

func Init() {
	userService = service.NewUserService(mongodb.UsersCollection, context.TODO())
	TokenService = service.NewTokenService(os.Getenv("JWT_SECRET"))
	UserController = controller.NewUserController(userService)
	AuthController = controller.NewAuthController(userService, TokenService)
}
