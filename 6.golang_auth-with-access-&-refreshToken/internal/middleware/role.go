package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireAuth checks that the user is logged in (role exists in context)
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(ctxRoleKey)
		if role == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}

// RequireRoles allows access only to the given roles (e.g. "user", "admin")
func RequireRoles(allowed ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(ctxRoleKey)

		if role == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		for _, r := range allowed {
			if role == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}

// RequiredAdmin is now just a shortcut using RequireRoles
func RequiredAdmin() gin.HandlerFunc {
	return RequireRoles("admin")
}
