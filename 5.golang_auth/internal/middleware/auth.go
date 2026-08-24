
package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// store ->auth data info ->gin context
const (
	ctxUserIDKey = "auth_userId"
	ctxRoleKey   = "auth_role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an authorization header"})
			return
		}

		// "Bearer " စာသားပါမပါ စစ်ဆေးပြီး token သီးသန့် ဖြတ်ထုတ်ခြင်း
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// Bearer မပါဘဲ လာရင် error ပြန်မည်
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format, must be 'Bearer <token>'"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Token ထဲက user ID နှင့် role ကို context ထဲ ထည့်သွင်းခြင်း
			c.Set(ctxUserIDKey, claims["sub"].(string))
			if role, ok := claims["role"].(string); ok {
				c.Set(ctxRoleKey, role)
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "failed to extract token claims"})
			return
		}

		c.Next()
	}
}

func GetUserID(c *gin.Context) string {
    return c.GetString(ctxUserIDKey)
}
