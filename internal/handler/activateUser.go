package handler

import (
	"go-auth/config"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Activate(c *gin.Context) {
	//parse link from param and convert to URL
	link := config.GetApiHost() + "/activate/" + c.Param("link")

	//fetch user by id and decode document
	result := mongodb.Users.FindOne(c, primitive.M{"activationLink": link})
	user := model.User{}
	err := result.Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found."})
		return
	}

	//update isActivated field to true
	_, err = mongodb.Users.UpdateOne(c, bson.M{"activationLink": link}, bson.M{"$set": bson.M{"isActivated": true}})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user cant be activated."})
		return
	}

	//send res
	c.JSON(http.StatusOK, gin.H{"message": "user activated."})
}
