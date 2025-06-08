package models

// City представляет информацию о городе России
type City struct {
	Address         string  `json:"address" csv:"address"`                   // Адрес одной строкой
	PostalCode      string  `json:"postal_code" csv:"postal_code"`           // Почтовый индекс
	Country         string  `json:"country" csv:"country"`                   // Страна
	FederalDistrict string  `json:"federal_district" csv:"federal_district"` // Федеральный округ
	RegionType      string  `json:"region_type" csv:"region_type"`           // Тип региона
	Region          string  `json:"region" csv:"region"`                     // Регион
	AreaType        string  `json:"area_type" csv:"area_type"`               // Тип района
	Area            string  `json:"area" csv:"area"`                         // Район
	CityType        string  `json:"city_type" csv:"city_type"`               // Тип города
	City            string  `json:"city" csv:"city"`                         // Город
	SettlementType  string  `json:"settlement_type" csv:"settlement_type"`   // Тип населенного пункта
	Settlement      string  `json:"settlement" csv:"settlement"`             // Населенный пункт
	KladrID         string  `json:"kladr_id" csv:"kladr_id"`                 // КЛАДР-код
	FiasID          string  `json:"fias_id" csv:"fias_id"`                   // ФИАС-код
	FiasLevel       int     `json:"fias_level" csv:"fias_level"`             // Уровень по ФИАС
	CapitalMarker   int     `json:"capital_marker" csv:"capital_marker"`     // Признак центра региона или района
	Okato           string  `json:"okato" csv:"okato"`                       // Код ОКАТО
	Oktmo           string  `json:"oktmo" csv:"oktmo"`                       // Код ОКТМО
	TaxOffice       string  `json:"tax_office" csv:"tax_office"`             // Код ИФНС
	Timezone        string  `json:"timezone" csv:"timezone"`                 // Часовой пояс
	GeoLat          float64 `json:"geo_lat" csv:"geo_lat"`                   // Широта
	GeoLon          float64 `json:"geo_lon" csv:"geo_lon"`                   // Долгота
	Population      int     `json:"population" csv:"population"`             // Население
	FoundationYear  int     `json:"foundation_year" csv:"foundation_year"`   // Год основания
}
