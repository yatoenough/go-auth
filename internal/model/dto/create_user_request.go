package dto

type CreateUserRequest struct {
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required,min=8,max=50"`
}
