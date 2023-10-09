package v1companystock

type BuyOrSellCompanyStocksRequest struct {
	CompanyStockID string `json:"company_stock_id"`
	BuyOrSell      string `json:"buy_or_sell"`
	Quantity       int    `json:"quantity"`
	GameNumber     int    `json:"game_number"`
}

type BuyCompanyStocksErrorMsg struct {
	Error string `json:"error"`
}
