package middleware

// Logger Middleware

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	} // TODO: Implement Logger middleware
}
