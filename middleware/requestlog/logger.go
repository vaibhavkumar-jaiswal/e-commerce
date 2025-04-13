// Package requestlog provides a Gin middleware for logging HTTP request details
// in plain text format using the logrus logger. Logs are saved to date-based log files.
package requestlog

import (
	"os"
	"path/filepath"
	"time"

	"e-commerce/utils/constants"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger returns a Gin middleware that logs each HTTP request in plain text format.
// The log includes method, path, status, latency, client IP, and request ID.
// Logs are written to a file named by the current date under the given logFilePath.
func Logger(logFilePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		status := 0
		clientIP := c.ClientIP()
		requestID, _ := c.Get("requestID")

		if err := os.MkdirAll(logFilePath, constants.DirPerm750); err != nil {
			logrus.WithError(err).Error("Failed to create log directory")
			return
		}

		fileName := "log-" + time.Now().Format(time.DateOnly) + ".log"
		fileLocation := filepath.Join(logFilePath, fileName)

		logFile, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, constants.FilePerm640) // #nosec G304
		if err != nil {
			logrus.WithError(err).Error("Failed to open log file")
			return
		}
		defer func() {
			err := logFile.Close()
			if err != nil {
				logrus.WithError(err).Error("Failed to close log file")
			}
		}()

		logger := logrus.New()
		logger.SetOutput(logFile)
		logger.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp: true,
		})

		c.Next()
		status = c.Writer.Status()
		latency := time.Since(start)

		logger.Infof("[LOG] %s %s %d %s %s | requestID: %v", method, path, status, latency, clientIP, requestID)
	}
}
