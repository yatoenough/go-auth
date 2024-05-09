package controller

import (
	"go-auth/internal/common/utils"
	"go-auth/internal/model/dto"
	"go-auth/internal/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService  service.UserService
	TokenService service.TokenService
}

func NewAuthController(userService service.UserService, tokenService service.TokenService) AuthController {
	return AuthController{
		UserService:  userService,
		TokenService: tokenService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var body dto.CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		dto.ApiResponse(c, http.StatusBadRequest, "Bad request.")
		return
	}

	_, err := ac.UserService.GetUserByEmail(&body.Email)
	if err == nil {
		dto.ApiResponse(c, http.StatusBadRequest, "User with e-mail `"+body.Email+"` already exists. Please login.")
		return
	}

	user, err := ac.UserService.CreateUser(&body)
	if err != nil {
		dto.ApiResponse(c, http.StatusBadGateway, "Unable to create user.")
		return
	}

	apiHost := os.Getenv("API_HOST")
	activationLink := apiHost + "/activate/" + user.ActivationCode
	msgBody := "Activate your account by following this link: " + activationLink

	err = utils.SendMail(user.Email, msgBody)
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to send mail.")
		return
	}

	token, err := ac.TokenService.CreateToken(*user)
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to generate authorization token.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": token})
}

func (ac *AuthController) Login(c *gin.Context) {
	var body dto.CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		dto.ApiResponse(c, http.StatusBadRequest, "Bad request.")
		return
	}

	user, err := ac.UserService.GetUserByEmail(&body.Email)
	if err != nil {
		dto.ApiResponse(c, http.StatusNotFound, "User with e-mail `"+body.Email+"` not found.")
		return
	}

	token, err := ac.TokenService.CreateToken(*user)
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to generate authorization token.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": token})
}
