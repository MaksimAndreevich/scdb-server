package repository

import (
	"github.com/scdb/server/internal/database"
	"github.com/scdb/server/internal/logger"
	"github.com/scdb/server/internal/models"
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

func GetStats() (models.TotalCountStats, []models.TotalCountByEducationType, error) {
	db := database.DB

	var totalCount models.TotalCountStats
	err := db.QueryRow(TotalCountQuery).Scan(
		&totalCount.TotalOrganizations,
		&totalCount.TotalDistricts,
		&totalCount.TotalRegions,
		&totalCount.TotalCities,
	)

	if err != nil {
		logger.Error("[GET STATS REPOSITORY] Ошибка при получении общего количества организаций ", err)
		return models.TotalCountStats{}, []models.TotalCountByEducationType{}, err
	}

	rows, err := db.Query(TotalCountByEducationTypeQuery)
	if err != nil {
		logger.Error("[GET STATS REPOSITORY] Ошибка при выполнении запроса статистики по типам образования ", err)
		return models.TotalCountStats{}, []models.TotalCountByEducationType{}, err
	}
	defer rows.Close() // Важно закрыть rows

	var totalCountByEducationType []models.TotalCountByEducationType
	for rows.Next() {
		var oneTypeStats models.TotalCountByEducationType
		err := rows.Scan(&oneTypeStats.EducationType, &oneTypeStats.TotalCount)
		if err != nil {
			logger.Error("[GET STATS REPOSITORY] Ошибка при сканировании строки статистики ", err)
			return models.TotalCountStats{}, []models.TotalCountByEducationType{}, err
		}
		totalCountByEducationType = append(totalCountByEducationType, oneTypeStats)
	}

	// Проверяем ошибки после итерации
	if err = rows.Err(); err != nil {
		logger.Error("[GET STATS REPOSITORY] Ошибка при итерации по строкам статистики ", err)
		return models.TotalCountStats{}, []models.TotalCountByEducationType{}, err
	}

	return totalCount, totalCountByEducationType, nil
}
