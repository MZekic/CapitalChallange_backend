package v1Companies

type GetCompanyInfoRequest struct {
	Ticker string `json:"ticker" db:"ticker"`
}
