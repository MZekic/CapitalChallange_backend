package v1companystock

import "capital-challenge-server/models"

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

type CompanyStockListQueryParams struct {
	Page     string
	PageSize string
}

type CompanyStockTransactionResponse struct {
	UserTransaction models.UserTransactions `json:"user_transaction"`
	CurrentBalance  float32 `json:"current_balance" db:"current_balance"`
}
