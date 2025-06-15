package models

// FilterParams представляет параметры фильтрации
type FilterParams struct {
	FederalDistrictID *int    `json:"federal_district_id,omitempty"`
	RegionID          *int    `json:"region_id,omitempty"`
	CityID            *string `json:"city_id,omitempty"`
	EducationTypeKey  *string `json:"education_type_key,omitempty"`
	Search            *string `json:"search,omitempty"`
}
