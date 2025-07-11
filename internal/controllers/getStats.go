package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scdb/server/internal/repository"
)

const TotalCountQuery = `
		SELECT 
        (SELECT COUNT(*) FROM education_organizations) as total_organizations,
        (SELECT COUNT(*) FROM federal_districts) as total_districts,
        (SELECT COUNT(*) FROM regions) as total_regions,
        (SELECT COUNT(*) FROM cities) as total_cities;`

const TotalCountByEducationTypeQuery = `
		SELECT 
			et.title as education_type,
			COUNT(eo.id) as organizations_count
		FROM education_types et
		LEFT JOIN education_organizations eo ON eo.fk_education_type_key = et.key
		GROUP BY et.title
		ORDER BY et.title;`

func GetStats(c *gin.Context) {
	total, totalByType, err := repository.GetStats()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении статистики"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":                   total,
		"total_by_education_type": totalByType,
	})

}
