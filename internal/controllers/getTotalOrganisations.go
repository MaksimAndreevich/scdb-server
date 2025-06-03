package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/scdb/server/internal/database"
	"gitlab.com/scdb/server/internal/logger"
)

const Query = `SELECT COUNT(*) AS total FROM education_organizations;`

func GetTotalOrganisations(c *gin.Context) {

	db := database.DB

	var total int
	err := db.QueryRow(Query).Scan(&total)

	if err != nil {
		logger.Error("Ошибка при получении общего числа огранизаций", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении организаций"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})

}
