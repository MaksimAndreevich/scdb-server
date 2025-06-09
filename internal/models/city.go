package models

import "time"

type City struct {
	FiasID            string    `json:"fias_id" db:"fias_id"`
	Address           string    `json:"address" db:"address"`
	PostalCode        string    `json:"postal_code" db:"postal_code"`
	Country           string    `json:"country" db:"country"`
	RegionID          int       `json:"region_id" db:"region_id"`
	FederalDistrictID int       `json:"federal_district_id" db:"federal_district_id"`
	RegionType        string    `json:"region_type" db:"region_type"`
	AreaType          string    `json:"area_type" db:"area_type"`
	Area              string    `json:"area" db:"area"`
	CityType          string    `json:"city_type" db:"city_type"`
	City              string    `json:"city" db:"city"`
	SettlementType    string    `json:"settlement_type" db:"settlement_type"`
	Settlement        string    `json:"settlement" db:"settlement"`
	KladrID           string    `json:"kladr_id" db:"kladr_id"`
	FiasLevel         int       `json:"fias_level" db:"fias_level"`
	CapitalMarker     int       `json:"capital_marker" db:"capital_marker"`
	OKATO             string    `json:"okato" db:"okato"`
	OKTMO             string    `json:"oktmo" db:"oktmo"`
	TaxOffice         string    `json:"tax_office" db:"tax_office"`
	Timezone          string    `json:"timezone" db:"timezone"`
	GeoLat            float64   `json:"geo_lat" db:"geo_lat"`
	GeoLon            float64   `json:"geo_lon" db:"geo_lon"`
	Population        int       `json:"population" db:"population"`
	FoundationYear    int       `json:"foundation_year" db:"foundation_year"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`

	// Связи
	Region          *Region                 `json:"region,omitempty" db:"-"`
	FederalDistrict *FederalDistrict        `json:"federal_district,omitempty" db:"-"`
	Organizations   []EducationOrganization `json:"organizations,omitempty" db:"-"`
}
