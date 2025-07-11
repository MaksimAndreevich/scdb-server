package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scdb/server/internal/config"
	"github.com/scdb/server/internal/repository"
)

// GetOrganizations обрабатывает GET /api/organizations?search=Борисовские пруды&federal_district_id=1&region_id=77&city_id=0c5b2444-70a0-4932-980c-b4dc0d3f02b5&education_type_key=general_education&page=1&per_page=20
func GetOrganizations(c *gin.Context) {
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))

	if page < 1 {
		page = 1
	}

	perPage, _ := strconv.Atoi(c.Query("per_page"))

	if perPage < 1 {
		perPage = 20
	}

	// В демо-режиме ограничиваем размер страницы
	if config.AppConfig.DemoMode && perPage > 50 {
		perPage = 50
	}

	// Получаем опциональные параметры фильтрации
	var federalDistrictID *int
	if fdID, err := strconv.Atoi(c.Query("federal_district_id")); err == nil {
		federalDistrictID = &fdID
	}

	var regionID *int
	if rID, err := strconv.Atoi(c.Query("region_id")); err == nil {
		regionID = &rID
	}

	var cityID *string
	if cID := c.Query("city_id"); cID != "" {
		cityID = &cID
	}

	var educationTypeKey *string
	if etKey := c.Query("education_type_key"); etKey != "" {
		educationTypeKey = &etKey
	}

	organizations, total, err := repository.GetOrganizations(
		search,
		federalDistrictID,
		regionID,
		cityID,
		educationTypeKey,
		perPage,
		(page-1)*perPage,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        organizations,
		"total":       total,
		"page":        page,
		"per_page":    perPage,
		"total_pages": (total + perPage - 1) / perPage,
	})
}
