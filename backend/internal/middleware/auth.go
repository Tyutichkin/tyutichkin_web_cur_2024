package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
	"strings"
)

var jwtKey = []byte("secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty auth token"})
			c.Abort()
			return
		}
		tokenString := strings.Split(tokenHeader, " ")[1]

		// Проверка токена
		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong auth token"})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
			c.Set("login", claims.Login)
			c.Set("isAdmin", claims.IsAdmin)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong auth token"})
			c.Abort()
			return
		}
	}
}
