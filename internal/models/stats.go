package models

type TotalCountStats struct {
	TotalOrganizations int `json:"total_organizations"`
	TotalDistricts     int `json:"total_districts"`
	TotalRegions       int `json:"total_regions"`
	TotalCities        int `json:"total_cities"`
}

type TotalCountByEducationType struct {
	EducationType string `json:"education_type"`
	TotalCount    int    `json:"total_count"`
}
