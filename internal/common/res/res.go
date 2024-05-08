package res

import (
	"time"

	"github.com/gin-gonic/gin"
)

// create new response with string message
func NewMessageResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"statusCode": statusCode, "message": message, "timestamp": time.Now()})
}

// create new response with struct body
func NewBodyResponse(c *gin.Context, statusCode int, body interface{}) {
	c.JSON(statusCode, gin.H{"statusCode": statusCode, "body": body, "timestamp": time.Now()})
}
