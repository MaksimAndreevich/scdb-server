package models

import "encoding/xml"

// Описание полей можно посмотреть в scructure-example.xml

type OpenData struct {
	XMLName      xml.Name      `xml:"OpenData" json:"-"`
	Certificates []Certificate `xml:"Certificates>Certificate" json:"certificates"`
}

type Certificate struct {
	XMLName                          xml.Name              `xml:"Certificate" json:"-"`
	ID                               string                `xml:"Id" json:"id"`
	IsFederal                        string                `xml:"IsFederal" json:"isFederal"` // 1 - федеральное, 0 - региональное
	StatusName                       string                `xml:"StatusName" json:"statusName"`
	TypeName                         string                `xml:"TypeName" json:"typeName"`
	RegionName                       string                `xml:"RegionName" json:"regionName"`
	RegionCode                       string                `xml:"RegionCode" json:"regionCode"`
	FederalDistrictName              string                `xml:"FederalDistrictName" json:"federalDistrictName"`
	FederalDistrictShortName         string                `xml:"FederalDistrictShortName" json:"federalDistrictShortName"`
	RegNumber                        string                `xml:"RegNumber" json:"regNumber"`
	SerialNumber                     string                `xml:"SerialNumber" json:"serialNumber"`
	FormNumber                       string                `xml:"FormNumber" json:"formNumber"`
	IssueDate                        string                `xml:"IssueDate" json:"issueDate"`
	EndDate                          string                `xml:"EndDate" json:"endDate"`
	ControlOrgan                     string                `xml:"ControlOrgan" json:"controlOrgan"`
	PostAddress                      string                `xml:"PostAddress" json:"postAddress"`
	EduOrgFullName                   string                `xml:"EduOrgFullName" json:"eduOrgFullName"`
	EduOrgShortName                  string                `xml:"EduOrgShortName" json:"eduOrgShortName"`
	EduOrgINN                        string                `xml:"EduOrgINN" json:"eduOrgINN"`
	EduOrgKPP                        string                `xml:"EduOrgKPP" json:"eduOrgKPP"`
	EduOrgOGRN                       string                `xml:"EduOrgOGRN" json:"eduOrgOGRN"`
	IndividualEntrepreneurLastName   string                `xml:"IndividualEntrepreneurLastName" json:"individualEntrepreneurLastName"`
	IndividualEntrepreneurFirstName  string                `xml:"IndividualEntrepreneurFirstName" json:"individualEntrepreneurFirstName"`
	IndividualEntrepreneurMiddleName string                `xml:"IndividualEntrepreneurMiddleName" json:"individualEntrepreneurMiddleName"`
	IndividualEntrepreneurAddress    string                `xml:"IndividualEntrepreneurAddress" json:"individualEntrepreneurAddress"`
	IndividualEntrepreneurEGRIP      string                `xml:"IndividualEntrepreneurEGRIP" json:"individualEntrepreneurEGRIP"`
	IndividualEntrepreneurINN        string                `xml:"IndividualEntrepreneurINN" json:"individualEntrepreneurINN"`
	ActualEducationOrganization      EducationOrganization `xml:"ActualEducationOrganization" json:"actualEducationOrganization"`
	Supplements                      []Supplement          `xml:"Supplements>Supplement" json:"supplements"`
	Decisions                        []Decision            `xml:"Decisions>Decision" json:"decisions"`
}

type EducationOrganization struct {
	XMLName                  xml.Name `xml:"ActualEducationOrganization" json:"-"`
	ID                       string   `xml:"Id" json:"id"`
	FullName                 string   `xml:"FullName" json:"fullName"`
	ShortName                string   `xml:"ShortName" json:"shortName"`
	HeadEduOrgId             string   `xml:"HeadEduOrgId" json:"headEduOrgId"`
	IsBranch                 string   `xml:"IsBranch" json:"isBranch"` // 1 - да, 0 - нет
	PostAddress              string   `xml:"PostAddress" json:"postAddress"`
	Phone                    string   `xml:"Phone" json:"phone"`
	Fax                      string   `xml:"Fax" json:"fax"`
	Email                    string   `xml:"Email" json:"email"`
	WebSite                  string   `xml:"WebSite" json:"webSite"`
	OGRN                     string   `xml:"OGRN" json:"ogrn"`
	INN                      string   `xml:"INN" json:"inn"`
	KPP                      string   `xml:"KPP" json:"kpp"`
	HeadPost                 string   `xml:"HeadPost" json:"headPost"`
	HeadName                 string   `xml:"HeadName" json:"headName"`
	FormName                 string   `xml:"FormName" json:"formName"`
	KindName                 string   `xml:"KindName" json:"kindName"`
	TypeName                 string   `xml:"TypeName" json:"typeName"`
	RegionName               string   `xml:"RegionName" json:"regionName"`
	FederalDistrictShortName string   `xml:"FederalDistrictShortName" json:"federalDistrictShortName"`
	FederalDistrictName      string   `xml:"FederalDistrictName" json:"federalDistrictName"`
}

type Supplement struct {
	XMLName                     xml.Name              `xml:"Supplement" json:"-"`
	ID                          string                `xml:"Id" json:"id"`
	StatusName                  string                `xml:"StatusName" json:"statusName"`
	StatusCode                  string                `xml:"StatusCode" json:"statusCode"`
	Number                      string                `xml:"Number" json:"number"`
	SerialNumber                string                `xml:"SerialNumber" json:"serialNumber"`
	FormNumber                  string                `xml:"FormNumber" json:"formNumber"`
	IssueDate                   string                `xml:"IssueDate" json:"issueDate"`
	IsForBranch                 string                `xml:"IsForBranch" json:"isForBranch"` // 1 - филиалу, 0 - головной
	Note                        string                `xml:"Note" json:"note"`
	EduOrgFullName              string                `xml:"EduOrgFullName" json:"eduOrgFullName"`
	EduOrgShortName             string                `xml:"EduOrgShortName" json:"eduOrgShortName"`
	EduOrgAddress               string                `xml:"EduOrgAddress" json:"eduOrgAddress"`
	EduOrgKPP                   string                `xml:"EduOrgKPP" json:"eduOrgKPP"`
	ActualEducationOrganization EducationOrganization `xml:"ActualEducationOrganization" json:"actualEducationOrganization"`
	EducationalPrograms         []EducationalProgram  `xml:"EducationalPrograms>EducationalProgram" json:"educationalPrograms"`
}

type EducationalProgram struct {
	XMLName            xml.Name `xml:"EducationalProgram" json:"-"`
	ID                 string   `xml:"Id" json:"id"`
	TypeName           string   `xml:"TypeName" json:"typeName"`
	EduLevelName       string   `xml:"EduLevelName" json:"eduLevelName"`
	ProgrammName       string   `xml:"ProgrammName" json:"programmName"`
	ProgrammCode       string   `xml:"ProgrammCode" json:"programmCode"`
	UGSName            string   `xml:"UGSName" json:"ugsName"`
	UGSCode            string   `xml:"UGSCode" json:"ugsCode"`
	EduNormativePeriod string   `xml:"EduNormativePeriod" json:"eduNormativePeriod"`
	Qualification      string   `xml:"Qualification" json:"qualification"`
	IsAccredited       string   `xml:"IsAccredited" json:"isAccredited"` // 0 - аккредитована, 1 - отказ
	IsCanceled         string   `xml:"IsCanceled" json:"isCanceled"`
	IsSuspended        string   `xml:"IsSuspended" json:"isSuspended"`
}

type Decision struct {
	XMLName             xml.Name `xml:"Decision" json:"-"`
	ID                  string   `xml:"Id" json:"id"`
	DecisionTypeName    string   `xml:"DecisionTypeName" json:"decisionTypeName"`
	OrderDocumentNumber string   `xml:"OrderDocumentNumber" json:"orderDocumentNumber"`
	OrderDocumentKind   string   `xml:"OrderDocumentKind" json:"orderDocumentKind"`
	DecisionDate        string   `xml:"DecisionDate" json:"decisionDate"`
}
