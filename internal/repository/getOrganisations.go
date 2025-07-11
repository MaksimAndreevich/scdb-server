package repository

import (
	"fmt"
	"strings"

	"github.com/scdb/server/internal/database"
	"github.com/scdb/server/internal/logger"
	"github.com/scdb/server/internal/models"
)

// GetOrganizations получает список организаций с фильтрацией и пагинацией
func GetOrganizations(
	search string,
	federalDistrictID *int,
	regionID *int,
	cityID *string,
	educationTypeKey *string,
	limit, offset int,
) ([]models.EducationOrganization, int, error) {
	db := database.DB

	// Базовый запрос
	baseQuery := `
        SELECT 
            id, full_name, short_name, head_edu_org_id, is_branch,
            post_address, phone, fax, email, web_site,
            ogrn, inn, kpp, head_post, head_name,
            form_name, kind_name, type_name, fk_city_id, fk_region_id,
            fk_federal_district_id, fk_education_type_key
        FROM education_organizations
        WHERE 1=1
    `

	// Параметры для запроса
	var params []interface{}
	var conditions []string
	paramIndex := 1

	// Добавляем поиск по всем полям
	if search != "" {
		searchCondition := fmt.Sprintf(`(
            LOWER(full_name) LIKE LOWER($%d::text) OR
            LOWER(short_name) LIKE LOWER($%d::text) OR
            LOWER(post_address) LIKE LOWER($%d::text)
        )`, paramIndex, paramIndex+1, paramIndex+2)

		conditions = append(conditions, searchCondition)
		searchPattern := "%" + search + "%"
		params = append(params, searchPattern, searchPattern, searchPattern)
		paramIndex += 3
	}

	// Добавляем фильтры
	if federalDistrictID != nil {
		conditions = append(conditions, fmt.Sprintf("fk_federal_district_id = $%d::integer", paramIndex))
		params = append(params, *federalDistrictID)
		paramIndex++
	}

	if regionID != nil {
		conditions = append(conditions, fmt.Sprintf("fk_region_id = $%d::integer", paramIndex))
		params = append(params, *regionID)
		paramIndex++
	}

	if cityID != nil {
		conditions = append(conditions, fmt.Sprintf("fk_city_id = $%d::text", paramIndex))
		params = append(params, *cityID)
		paramIndex++
	}

	if educationTypeKey != nil {
		conditions = append(conditions, fmt.Sprintf("fk_education_type_key = $%d::text", paramIndex))
		params = append(params, *educationTypeKey)
		paramIndex++
	}

	// Добавляем условия в запрос
	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Сначала получаем общее количество
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) as count_query", baseQuery)
	var total int
	err := db.QueryRow(countQuery, params...).Scan(&total)
	if err != nil {
		logger.Error("Ошибка при подсчете общего количества организаций", err)
		return nil, 0, err
	}

	// Затем добавляем сортировку и пагинацию для основного запроса
	baseQuery += " ORDER BY full_name"
	baseQuery += fmt.Sprintf(" LIMIT $%d::integer OFFSET $%d::integer", paramIndex, paramIndex+1)

	// Создаем новый слайс параметров для основного запроса
	mainParams := make([]interface{}, len(params))
	copy(mainParams, params)
	mainParams = append(mainParams, limit, offset)

	// Выполняем основной запрос
	rows, err := db.Query(baseQuery, mainParams...)
	if err != nil {
		logger.Error("Ошибка при получении организаций", err)
		return nil, 0, err
	}

	defer rows.Close()

	var organizations []models.EducationOrganization
	for rows.Next() {
		var org models.EducationOrganization
		err := rows.Scan(
			&org.ID,
			&org.FullName,
			&org.ShortName,
			&org.HeadEduOrgID,
			&org.IsBranch,
			&org.PostAddress,
			&org.Phone,
			&org.Fax,
			&org.Email,
			&org.WebSite,
			&org.OGRN,
			&org.INN,
			&org.KPP,
			&org.HeadPost,
			&org.HeadName,
			&org.FormName,
			&org.KindName,
			&org.TypeName,
			&org.CityID,
			&org.RegionID,
			&org.FederalDistrictID,
			&org.EducationTypeKey,
		)
		if err != nil {
			logger.Error("Ошибка при чтении строки результата", err)
			return nil, 0, err
		}
		organizations = append(organizations, org)
	}

	if err = rows.Err(); err != nil {
		logger.Error("Ошибка после итерации по строкам", err)
		return nil, 0, err
	}

	return organizations, total, nil
}
