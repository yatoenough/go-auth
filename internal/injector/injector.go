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
	mailService    service.MailService
	AuthMiddleware middleware.AuthMiddleware
	UserController controller.UserController
	AuthController controller.AuthController
)

func InitDependencies() {
	apiHost := os.Getenv("API_HOST")
	senderMail := os.Getenv("SENDER_MAIL")
	appPassword := os.Getenv("APP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	accessSecret := os.Getenv("JWT_ACCESS_SECRET")
	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")

	userService = service.NewUserService(mongodb.UsersCollection, context.TODO())
	tokenService = service.NewTokenService(accessSecret, refreshSecret)
	mailService = service.NewMailService(apiHost, senderMail, appPassword, smtpHost, smtpPort)

	AuthMiddleware = middleware.NewAuthMiddleware(tokenService)

	UserController = controller.NewUserController(userService)
	AuthController = controller.NewAuthController(userService, tokenService, mailService)
}
