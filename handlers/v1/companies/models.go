package v1Companies

import "capital-challenge-server/models"

type GetCompanyInfoRequest struct {
	Ticker string `json:"ticker" db:"ticker"`
}

type GetCompanyInfoResponse struct {
	Company      models.Companies    `json:"company"`
	CompanyStock models.CompanyStock `json:"company_stock"`
}
