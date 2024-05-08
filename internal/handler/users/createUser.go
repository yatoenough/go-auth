package users

import (
	"go-auth/internal/common/res"
	"go-auth/internal/common/utils"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	//get body from req
	var body model.CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		res.NewMessageResponse(c, http.StatusBadRequest, "bad request")
		return
	}
	//create user entity
	user := model.NewUser(body)

	//insert user into db
	_, err := mongodb.Users.InsertOne(c, user)
	if err != nil {
		res.NewMessageResponse(c, http.StatusInternalServerError, "Unable to create user")
		return
	}

	//send activation link to email
	utils.SendMail(user.Email, user.ActivationCode)

	//send res
	res.NewBodyResponse(c, http.StatusCreated, user)
}
