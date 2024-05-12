package injector

import (
	"context"
	"go-auth/internal/controller"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/middleware"
	"go-auth/internal/service"
	"os"
)

var (
	userService    service.UserService
	tokenService   service.TokenService
	AuthMiddleware middleware.AuthMiddleware
	UserController controller.UserController
	AuthController controller.AuthController
)

func InitDependencies() {
	userService = service.NewUserService(mongodb.UsersCollection, context.TODO())
	tokenService = service.NewTokenService(os.Getenv("JWT_SECRET"))

	AuthMiddleware = middleware.NewAuthMiddleware(tokenService)

	UserController = controller.NewUserController(userService)
	AuthController = controller.NewAuthController(userService, tokenService)
}
