package users

import (
	"go-auth/internal/common/res"
	"go-auth/internal/database/mongodb"
	"go-auth/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *gin.Context) {
	//fetch users from mongodb
	cursor, err := mongodb.Users.Find(c, bson.M{})
	if err != nil {
		res.NewMessageResponse(c, http.StatusInternalServerError, "Unable to fetch users.")
		return
	}

	//decode all documents into result
	var users []model.User
	if err = cursor.All(c, &users); err != nil {
		res.NewMessageResponse(c, http.StatusInternalServerError, "Unable to fetch users.")
		return
	}

	//send res
	res.NewBodyResponse(c, http.StatusOK, users)
}
