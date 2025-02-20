package middleware

import (
	"net/http"
	"strings"

	"github.com/1991-bishnu/loan-service/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(conf *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" || tokenParts[1] != conf.Auth.Token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
