package dto

import (
	"time"

	"github.com/gin-gonic/gin"
)

func ApiResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"statusCode": statusCode, "message": message, "timestamp": time.Now()})
}
