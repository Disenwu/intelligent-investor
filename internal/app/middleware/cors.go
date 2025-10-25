package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
			c.Header("Allow", "POST, GET, OPTIONS, PUT, DELETE")
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
