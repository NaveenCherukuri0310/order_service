package middleware

import (
	"fmt"
	"order_service/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		c.Next()
		duration := time.Since(start)

		message := fmt.Sprintf(
			"%s %s | Status=%d | Time=%v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)

		// write to file
		logger.Log.Println(message)
	}
}
