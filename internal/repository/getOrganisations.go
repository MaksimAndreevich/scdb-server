package repository

import (
	"gitlab.com/scdb/server/internal/database"
	"gitlab.com/scdb/server/internal/logger"
	"gitlab.com/scdb/server/internal/models"
)

const Query = `
		SELECT 
			id, full_name, short_name, head_edu_org_id, is_branch,
			post_address, phone, fax, email, web_site,
			ogrn, inn, kpp, head_post, head_name,
			form_name, kind_name, type_name, region_name,
			federal_district_short_name, federal_district_name
		FROM education_organizations
		WHERE
			LOWER(full_name) LIKE LOWER($1) OR
			LOWER(short_name) LIKE LOWER($1) OR
			LOWER(region_name) LIKE LOWER($1) OR
			LOWER(post_address) LIKE LOWER($1) OR
			LOWER(head_name) LIKE LOWER($1) OR
			LOWER(federal_district_name) LIKE LOWER($1)
		ORDER BY full_name
		LIMIT $2 OFFSET $3;
	`

func GetOrganizations(search string, limit, offset int) ([]models.EducationOrganization, error) {
	db := database.DB

	rows, err := db.Query(Query, "%"+search+"%", limit, offset)

	if err != nil {
		logger.Error("Ошибка при получении организации ", err)
	}

	defer rows.Close()

	var result []models.EducationOrganization

	for rows.Next() {
		var org models.EducationOrganization

		err := rows.Scan(
			&org.ID,
			&org.FullName,
			&org.ShortName,
			&org.HeadEduOrgId,
			org.IsBranch,
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
			&org.RegionName,
			&org.FederalDistrictShortName,
			&org.FederalDistrictName,
		)
		if err != nil {
			logger.Error("Ошибка при чтении строки результата ", err)
			return nil, err
		}
		result = append(result, org)
	}

	// Проверяем ошибки после цикла
	if err = rows.Err(); err != nil {
		logger.Error("Ошибка после итерации по строкам ", err)
		return nil, err
	}

	return result, err
}
