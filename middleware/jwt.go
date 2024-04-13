package middleware

// JWT Middleware

import (
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	} // TODO: Implement JWT middleware
}
