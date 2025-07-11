package repository

import (
	"database/sql"

	"scdb-server/internal/database"
	"scdb-server/internal/logger"
	"scdb-server/internal/models"
)

const getByIDQuery = `
	SELECT 
		id, full_name, short_name, head_edu_org_id, is_branch,
		post_address, phone, fax, email, web_site,
		ogrn, inn, kpp, head_post, head_name,
		form_name, kind_name, type_name, fk_city_id, fk_region_id,
		fk_federal_district_id, fk_education_type_key
	FROM education_organizations
	WHERE id = $1;
`

func GetOrganisationById(id string) (*models.EducationOrganization, error) {
	db := database.DB

	var org models.EducationOrganization

	err := db.QueryRow(getByIDQuery, id).Scan(
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
		if err == sql.ErrNoRows {
			logger.Warning("Организация не найдена по ID: ", id)
			return nil, nil
		}

		logger.Error("Ошибка при получении организации по ID: ", err)
		return nil, err
	}

	return &org, nil
}
