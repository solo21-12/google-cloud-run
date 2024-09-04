package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		log.Printf("Received request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()

		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		log.Printf("Response: %d %v", statusCode, duration)
	}
}
