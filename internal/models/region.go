package models

import "time"

type Region struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Label       string  `json:"label" db:"label"`
	Type        string  `json:"type" db:"type"`
	TypeShort   string  `json:"type_short" db:"type_short"`
	ContentType string  `json:"content_type" db:"content_type"`
	RegionID    string  `json:"region_id" db:"region_id"`
	OKATO       string  `json:"okato" db:"okato"`
	OKTMO       string  `json:"oktmo" db:"oktmo"`
	GUID        string  `json:"guid" db:"guid"`
	Code        int     `json:"code" db:"code"`
	ISO3166_2   string  `json:"iso_3166_2" db:"iso_3166_2"`
	Population  int     `json:"population" db:"population"`
	YearFounded int     `json:"year_founded" db:"year_founded"`
	Area        float64 `json:"area" db:"area"`
	Fullname    string  `json:"fullname" db:"fullname"`
	NameEn      string  `json:"name_en" db:"name_en"`

	FederalDistrictID int    `json:"federal_district_id" db:"federal_district_id"`
	Cities            []City `json:"cities,omitempty" db:"-"`

	// Падежи названия
	NamecaseNominative    string `json:"namecase_nominative" db:"namecase_nominative"`
	NamecaseGenitive      string `json:"namecase_genitive" db:"namecase_genitive"`
	NamecaseDative        string `json:"namecase_dative" db:"namecase_dative"`
	NamecaseAccusative    string `json:"namecase_accusative" db:"namecase_accusative"`
	NamecaseAblative      string `json:"namecase_ablative" db:"namecase_ablative"`
	NamecasePrepositional string `json:"namecase_prepositional" db:"namecase_prepositional"`
	NamecaseLocative      string `json:"namecase_locative" db:"namecase_locative"`

	// Столица региона
	CapitalName        string `json:"capital_name" db:"capital_name"`
	CapitalLabel       string `json:"capital_label" db:"capital_label"`
	CapitalID          string `json:"capital_id" db:"capital_id"`
	CapitalOKATO       string `json:"capital_okato" db:"capital_okato"`
	CapitalOKTMO       string `json:"capital_oktmo" db:"capital_oktmo"`
	CapitalContentType string `json:"capital_content_type" db:"capital_content_type"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
