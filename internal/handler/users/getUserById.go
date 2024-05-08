package users

import (
	"go-auth/internal/database/mongodb"
	"go-auth/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserById(c *gin.Context) {
	//parse id from param and convert to ObjectID
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	//fetch user by id and decode document
	result := mongodb.Users.FindOne(c, primitive.M{"_id": _id})
	user := model.User{}
	err = result.Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found."})
		return
	}

	//send res
	c.JSON(http.StatusOK, user)
}
