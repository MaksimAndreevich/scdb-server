package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repo "gitlab.com/scdb/database/repository"
)

func GetOrganisationById(c *gin.Context) {
	id := c.Param("id")

	org, err := repo.GetOrganisationById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении организации по ID"})
		return
	}

	if org == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Организация не найдена"})
		return
	}

	c.JSON(http.StatusOK, org)
}
