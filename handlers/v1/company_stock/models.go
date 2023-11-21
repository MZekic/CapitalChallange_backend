package v1companystock

type BuyOrSellCompanyStocksRequest struct {
	CompanyStockID string `json:"company_stock_id"`
	BuyOrSell      string `json:"buy_or_sell"`
	Quantity       int    `json:"quantity"`
}

type BuyCompanyStocksErrorMsg struct {
	Error string `json:"error"`
}

type BuyCompanyStockRequest struct {
	CompanyStockID string `json:"company_stock_id"`
	Quantity       int    `json:"quantity"`
}

type SellCompanyStockRequest struct {
	CompanyStockID string `json:"company_stock_id"`
	Quantity       int    `json:"quantity"`
}
