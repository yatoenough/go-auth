package middleware

import (
	"go-auth/internal/common/utils"
	"go-auth/internal/injector"
	"go-auth/internal/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	jwtToken, err := utils.ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		dto.ApiResponse(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	token, err := injector.TokenService.VerifyToken(jwtToken)
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