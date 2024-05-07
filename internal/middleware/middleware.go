package middleware

import (
	"fetch/management-console/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Param("apikey")

		if key != config.API_KEY {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
