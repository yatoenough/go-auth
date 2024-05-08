package users

import (
	"go-auth/internal/common/res"
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
		res.NewMessageResponse(c, http.StatusBadRequest, "Invalid ID provided.")
		return
	}

	//delete user
	result, _ := mongodb.Users.DeleteOne(c, bson.M{"_id": _id})

	//check if user was deleted
	if result.DeletedCount == 0 {
		res.NewMessageResponse(c, http.StatusNotFound, "User with ID `"+id+"` not found.")
		return
	}

	//send res
	res.NewMessageResponse(c, http.StatusOK, "Deleted successfully!")
}
