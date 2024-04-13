package middleware

// CORS middleware
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// handle preflight request
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// handle request
		c.Next()
	}
}
