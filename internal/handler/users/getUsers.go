package users

import (
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch users"})
		return
	}

	//decode all documents into result
	var users []model.User
	if err = cursor.All(c, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch users"})
		return
	}

	//send res
	c.JSON(http.StatusOK, users)
}
