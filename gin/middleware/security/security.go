package security

import "github.com/gin-gonic/gin"

func SecurityHttpHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Next()
		c.Writer.Header().Add("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Add("X-Xss-Protection", "1; mode=block")
	}
}
