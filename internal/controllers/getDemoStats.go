package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scdb/server/internal/middleware"
)

func GetDemoStats(c *gin.Context) {
	stats := middleware.GetDemoStats()

	c.JSON(http.StatusOK, gin.H{
		"demo_stats": stats,
		"message":    "Демо-версия образовательной базы данных",
		"note":       "Ограничение: 50 запросов на образ. Обратитесь к разработчику для получения нового доступа.",
	})

}
