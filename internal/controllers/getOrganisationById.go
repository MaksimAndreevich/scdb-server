package controllers

import (
	"net/http"

	"scdb-server/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetOrganisationById(c *gin.Context) {
	id := c.Param("id")

	org, err := repository.GetOrganisationById(id)
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
