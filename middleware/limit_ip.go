package middleware

// LimitIP Middleware

import (
	"github.com/gin-gonic/gin"
)

func LimitIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	} // TODO: Implement LimitIP middleware
}
