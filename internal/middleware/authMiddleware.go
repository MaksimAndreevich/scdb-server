package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		if isAuthenticated(c) {
			c.Next()
			return
		}

		// Если не авторизован возвращаем ошибку
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Не авторизованный"})
	}
}

func isAuthenticated(c *gin.Context) bool {
	// TODO: проверяем авторизацию пользователя
	return true
}
