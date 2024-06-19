package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	logger := logrus.New()

	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()
		// Process request
		c.Next()
		// Stop timer
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		// Request method
		reqMethod := c.Request.Method
		// Request URI
		reqURI := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Client IP
		clientIP := c.ClientIP()

		// Log format
		entry := logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURI,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			entry.Info("Request processed")
		}
	}
}
