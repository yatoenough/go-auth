package controller

import (
	"go-auth/internal/model/dto"
	"go-auth/internal/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService  service.UserService
	tokenService service.TokenService
	mailService  service.MailService
}

func NewAuthController(
	userService service.UserService,
	tokenService service.TokenService,
	mailService service.MailService,
) AuthController {
	return AuthController{
		userService:  userService,
		tokenService: tokenService,
		mailService:  mailService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var body dto.CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		dto.ApiResponse(c, http.StatusBadRequest, "Bad request.")
		return
	}

	_, err := ac.userService.GetUserByEmail(&body.Email)
	if err == nil {
		dto.ApiResponse(c, http.StatusBadRequest, "User with e-mail `"+body.Email+"` already exists. Please login.")
		return
	}

	user, err := ac.userService.CreateUser(&body)
	if err != nil {
		dto.ApiResponse(c, http.StatusBadGateway, "Unable to create user.")
		return
	}

	apiHost := os.Getenv("API_HOST")
	activationLink := apiHost + "/activate/" + user.ActivationCode
	msgBody := "Activate your account by following this link: " + activationLink

	err = ac.mailService.SendActivationMail(user.Email, msgBody)
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to send mail.")
		return
	}

	tokens, err := ac.tokenService.GenerateTokens(*user)
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to generate authorization token.")
		return
	}

	c.SetCookie("refresh_token", tokens["refresh_token"], 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"accessToken": tokens["access_token"]})
}

func (ac *AuthController) Login(c *gin.Context) {
	var body dto.CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		dto.ApiResponse(c, http.StatusBadRequest, "Bad request.")
		return
	}

	user, err := ac.userService.GetUserByEmail(&body.Email)
	if err != nil {
		dto.ApiResponse(c, http.StatusNotFound, "User with e-mail `"+body.Email+"` not found.")
		return
	}

	tokens, err := ac.tokenService.GenerateTokens(*user)
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to generate authorization token.")
		return
	}

	c.SetCookie("refresh_token", tokens["refresh_token"], 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"accessToken": tokens["access_token"]})
}
