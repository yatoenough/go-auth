package service

import (
	"context"
	"errors"
	"go-auth/internal/model"
	"go-auth/internal/model/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(*dto.CreateUserRequest) (*model.User, error)
	ActivateUser(code *string) error
	GetUserById(*string) (*model.User, error)
	GetUserByCode(*string) (*model.User, error)
	GetAll() ([]model.User, error)
	UpdateUser(*model.User) error
	DeleteUser(*string) error
}

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(body *dto.CreateUserRequest) (*model.User, error) {
	user := model.NewUser(body)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.User{}, err
	}
	user.Password = string(hash)
	_, err = u.userCollection.InsertOne(u.ctx, user)
	return &user, err
}

func (u *UserServiceImpl) GetUserById(id *string) (*model.User, error) {
	var user *model.User
	_id, _ := primitive.ObjectIDFromHex(*id)
	query := bson.D{bson.E{Key: "_id", Value: _id}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetUserByCode(code *string) (*model.User, error) {
	var user *model.User
	query := bson.D{bson.E{Key: "activationCode", Value: code}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]model.User, error) {
	cursor, err := u.userCollection.Find(u.ctx, bson.M{})
	users := []model.User{}
	if err != nil {
		return users, err
	}
	err = cursor.All(u.ctx, &users)

	return users, err
}

func (u *UserServiceImpl) UpdateUser(*model.User) error {
	return nil
}

func (u *UserServiceImpl) ActivateUser(code *string) error {
	filter := bson.D{bson.E{Key: "activationCode", Value: code}}
	update := bson.D{bson.E{Key: "$set", Value: bson.M{"isActivated": true}}}
	result, _ := u.userCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}

	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	_id, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "_id", Value: _id}}
	result, _ := u.userCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
