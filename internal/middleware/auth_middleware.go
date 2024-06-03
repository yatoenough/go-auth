package middleware

import (
	"errors"
	"go-auth/internal/model/dto"
	"go-auth/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware interface {
	Guard(*gin.Context)
}

type authMiddlewareImpl struct {
	tokenService service.TokenService
}

func NewAuthMiddleware(tokenService service.TokenService) AuthMiddleware {
	return &authMiddlewareImpl{
		tokenService: tokenService,
	}
}

func (am *authMiddlewareImpl) Guard(c *gin.Context) {
	jwtToken, err := am.extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		dto.ApiResponse(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	token, err := am.tokenService.VerifyToken(jwtToken)
	if err != nil {
		dto.ApiResponse(c, http.StatusBadRequest, "Bad JWT Token.")
		c.Abort()
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		dto.ApiResponse(c, http.StatusInternalServerError, "Unable to parse claims.")
		c.Abort()
		return
	}

	c.Set("token_data", claims)
	c.Next()
}

func (am *authMiddlewareImpl) extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
