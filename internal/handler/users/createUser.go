package users

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	//create user entity
	user := model.NewUser(body)

	//insert user into db
	_, err := mongodb.Users.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user"})
		return
	}

	//send activation link to email
	utils.SendMail(user.Email, user.ActivationCode)

	//send res
	c.JSON(http.StatusCreated, user)
}
