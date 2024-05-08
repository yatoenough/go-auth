package users

import (
	"go-auth/internal/common/res"
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
		res.NewMessageResponse(c, http.StatusBadRequest, "Invalid ID provided.")
		return
	}

	//fetch user by id and decode document
	result := mongodb.Users.FindOne(c, primitive.M{"_id": _id})
	user := model.User{}
	err = result.Decode(&user)
	if err != nil {
		res.NewMessageResponse(c, http.StatusNotFound, "User with ID `"+id+"` not found.")
		return
	}

	//send res
	res.NewBodyResponse(c, http.StatusOK, user)
}
