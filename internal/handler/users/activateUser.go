package users

import (
	"go-auth/internal/database/mongodb"
	"go-auth/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ActivateUser(c *gin.Context) {
	//parse link from param and convert to URL
	code := c.Param("code")

	//fetch user by id and decode document
	result := mongodb.Users.FindOne(c, primitive.M{"activationCode": code})
	user := model.User{}
	err := result.Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found."})
		return
	}

	//update isActivated field to true
	_, err = mongodb.Users.UpdateOne(c, bson.M{"activationCode": code}, bson.M{"$set": bson.M{"isActivated": true}})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user cant be activated."})
		return
	}

	//send res
	c.JSON(http.StatusOK, gin.H{"message": "user activated."})
}
