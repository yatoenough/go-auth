package users

import (
	"go-auth/internal/database/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(c *gin.Context) {
	//parse id from param and convert to ObjectID
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	//delete user
	res, _ := mongodb.Users.DeleteOne(c, bson.M{"_id": _id})

	//check if user was deleted
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found."})
		return
	}

	//send res
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
