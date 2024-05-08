package model

import (
	"go-auth/config"
	"go-auth/internal/common/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// define user model
type User struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email          string             `bson:"email" json:"email"`
	Password       string             `bson:"password" json:"password"`
	IsActivated    bool               `bson:"isActivated" json:"isActivated"`
	ActivationLink string             `bson:"activationLink" json:"activationLink"`
}

type CreateUserRequest struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

// method to create user entity
func NewUser(body CreateUserRequest) User {
	user := User{
		Id:             primitive.NewObjectID(),
		Email:          body.Email,
		Password:       body.Password,
		IsActivated:    false,
		ActivationLink: config.GetApiHost() + "/activate/" + utils.GenerateRandomString(10),
	}
	return user
}
