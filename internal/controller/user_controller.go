package controller

import (
	"go-auth/internal/model/dto"
	"go-auth/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) ActivateUser(c *gin.Context) {
	code := c.Param("code")

	user, err := uc.UserService.GetUserByCode(&code)
	if err != nil {
		dto.ApiResponse(c, http.StatusNotFound, "User not found.")
		return
	}

	err = uc.UserService.ActivateUser(&user.ActivationCode)
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to activate user.")
		return
	}

	//send res
	dto.ApiResponse(c, http.StatusOK, "User activated successfully!")
}

func (uc *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.UserService.GetUserById((&id))
	if err != nil {
		dto.ApiResponse(c, http.StatusNotFound, "User not found.")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to fetch users.")
		return
	}
	c.JSON(http.StatusOK, users)

}

func (uc *UserController) UpdateUser(c *gin.Context) {

}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := uc.UserService.DeleteUser(&id)
	if err != nil {
		dto.ApiResponse(c, http.StatusInternalServerError, "User not found.")
		return
	}
	dto.ApiResponse(c, http.StatusOK, "User deleted.")
}
