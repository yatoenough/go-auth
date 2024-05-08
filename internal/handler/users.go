package handler

import (
	"go-auth/internal/database/mongodb"
	"go-auth/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *gin.Context) {
	var body model.CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	user := model.NewUser(body)

	_, err := mongodb.Users.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	cursor, err := mongodb.Users.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch users"})
		return
	}

	var users []model.User
	if err = cursor.All(c, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	result := mongodb.Users.FindOne(c, primitive.M{"_id": _id})
	user := model.User{}
	err = result.Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found."})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, _ := mongodb.Users.DeleteOne(c, bson.M{"_id": _id})

	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
