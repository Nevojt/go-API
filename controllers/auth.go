package controllers

import (
	"api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Отримуємо токен з заголовка Authorization
		authHeder := c.GetHeader("Authorization")
		if authHeder == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication is required"})
			c.Abort()
			return
		}
		//Видаляємо префікс "Bearer "
		tokenString := strings.TrimPrefix(authHeder, "Bearer ")
		if tokenString == authHeder {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must start with Bearer"})
			c.Abort()
			return
		}
		// Перевіряємо валідність токену
		claims, err := util.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		// Зберігаємо інформацію про користувача в контексті
		c.Set("user_id", claims.ID)
		c.Set("email", claims.Email)

		c.Next() // Викликаємо наступний обробник запиту
	}
}
