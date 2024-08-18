package Middleware

import (
	"Golang/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func unauthorizedResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "You must be authorized to perform this operation",
	})
}

func RequiresAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("auth-token")
		if token == "" {
			unauthorizedResponse(c)
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			unauthorizedResponse(c)
			return
		}

		c.Set("role", claims.Role)
		c.Next()
	}
}

func RequiresRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			unauthorizedResponse(c)
			return
		}

		if role != requiredRole {
			unauthorizedResponse(c)
		}

		return
	}
}
