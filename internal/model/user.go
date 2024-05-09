package model

import (
	"go-auth/internal/common/utils"
	"go-auth/internal/model/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// define user model
type User struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email          string             `bson:"email" json:"email"`
	Password       string             `bson:"password" json:"-"`
	IsActivated    bool               `bson:"isActivated" json:"isActivated"`
	ActivationCode string             `bson:"activationCode" json:"activationCode"`
}

func NewUser(body *dto.CreateUserRequest) User {
	user := User{
		Id:             primitive.NewObjectID(),
		Email:          body.Email,
		Password:       body.Password,
		IsActivated:    false,
		ActivationCode: utils.GenerateRandomString(10),
	}
	return user
}
