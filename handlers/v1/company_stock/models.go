package v1companystock

type BuyCompanyStocksRequest struct {
	CompanyStockID string `json:"company_stock_id"`
	Quantity       int    `json:"quantity"`
	GameNumber     int    `json:"game_number"`
}

type BuyCompanyStocksErrorMsg struct {
	Error string `json:"error"`
}
