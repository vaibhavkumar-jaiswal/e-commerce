// Package compression provides middleware to enable gzip compression for HTTP responses.
package compression

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// Compression returns a Gin middleware handler that applies gzip compression
// to all HTTP responses using the default compression level.
func Compression() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}
