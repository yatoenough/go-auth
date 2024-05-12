package service

import (
	"errors"
	"go-auth/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	CreateToken(model.User) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type tokenServiceImpl struct {
	secretKey []byte
}

func NewTokenService(secret string) TokenService {
	return &tokenServiceImpl{
		secretKey: []byte(secret),
	}
}

func (ts *tokenServiceImpl) CreateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   user.Id,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		})

	tokenString, err := token.SignedString(ts.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ts *tokenServiceImpl) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ts.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
