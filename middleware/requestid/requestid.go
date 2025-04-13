// Package requestid provides middleware to manage unique request IDs for tracing and logging.
package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID is a Gin middleware that ensures every incoming request has a unique ID.
// If the "X-Request-ID" header is not present, it generates a new UUID.
// The request ID is stored in the context and set in the response header.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("requestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}
