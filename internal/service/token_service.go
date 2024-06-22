package service

import (
	"errors"
	"go-auth/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateTokens(model.User) (map[string]string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type tokenServiceImpl struct {
	accessSecretKey  []byte
	refreshSecretKey []byte
}

func NewTokenService(accessSecret, refreshSecret string) TokenService {
	return &tokenServiceImpl{
		accessSecretKey:  []byte(accessSecret),
		refreshSecretKey: []byte(refreshSecret),
	}
}

func (ts *tokenServiceImpl) GenerateTokens(user model.User) (map[string]string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   user.Id,
			"email": user.Email,
			"exp":   time.Now().Add(time.Minute * 30).Unix(),
		})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   user.Id,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

	accessTokenString, err := accessToken.SignedString(ts.accessSecretKey)
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := refreshToken.SignedString(ts.accessSecretKey)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	}, nil
}

func (ts *tokenServiceImpl) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ts.accessSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
