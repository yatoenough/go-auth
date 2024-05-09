package service

import (
	"errors"
	"go-auth/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

type TokenService interface {
	CreateToken(model.User) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type TokenServiceImpl struct {
	secretKey []byte
}

func NewTokenService(secret string) TokenService {
	return &TokenServiceImpl{
		secretKey: []byte(secret),
	}
}

func (ts *TokenServiceImpl) CreateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   user.Id,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ts *TokenServiceImpl) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
