package middleware

import (
	"github.com/TakuroBreath/personal-finance-manager/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"

	"strings"
)

func AuthMiddleware(jwtService *service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Удаляем префикс "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Валидируем и парсим токен
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Извлекаем ID пользователя из claims
		userID, ok := claims["userID"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		// Сохраняем ID пользователя в контексте
		c.Set("userID", int(userID))
		c.Next()
	}
}
