package requestlog

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(logFilePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		fileName := "log-" + time.Now().Format(time.DateOnly) + ".log"
		currentDir, err := os.Getwd()
		if err != nil {
			log.Println("Error getting current working directory:", err)
			return
		}
		fileLocation := filepath.Join(currentDir, logFilePath, fileName)

		c.Next()

		latency := time.Since(start)

		f, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Error opening log file:", err)
			return
		}

		defer f.Close()

		log.SetOutput(f)
		log.Printf("[LOG] %s %s %d %s %s\n", method, path, c.Writer.Status(), latency, c.ClientIP())
	}
}
