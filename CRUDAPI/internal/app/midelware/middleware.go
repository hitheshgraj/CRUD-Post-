package middleware

import (
	"CRUDAPI/internal/app/confi"

	"github.com/gin-gonic/gin"
)

func BasicAuth(appconfig *confi.AppConfig) gin.HandlerFunc {
	// Create the Gin middleware handler for BasicAuth
	authMiddleware := gin.BasicAuth(gin.Accounts{
		"hithesh": "1234",
	})

	return func(c *gin.Context) {
		// Execute the BasicAuth middleware
		authMiddleware(c)

		// Continue to the next middleware or route handler
		c.Next()
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
