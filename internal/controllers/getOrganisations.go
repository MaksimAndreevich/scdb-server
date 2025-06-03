package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repo "gitlab.com/scdb/database/repository"
)

func GetOrganisations(c *gin.Context) {

	// Получаем query-параметры
	search := c.DefaultQuery("search", "")
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	// Конвертируем limit и offset в int
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	// Получаем данные из репозитория
	orgs, err := repo.GetOrganisations(search, limit, offset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении организаций"})
		return
	}

	// Отдаем json ответ
	c.JSON(http.StatusOK, gin.H{
		"data":   orgs,
		"search": search,
		"limit":  limit,
		"offset": offset,
	})
}
