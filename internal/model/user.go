package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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

func NewUser(body CreateUserRequest) User {
	user := User{
		Id:             primitive.NewObjectID(),
		Email:          body.Email,
		Password:       body.Password,
		IsActivated:    false,
		ActivationLink: "test",
	}
	return user
}
