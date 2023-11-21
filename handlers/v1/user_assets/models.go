package v1userassets

import "capital-challenge-server/models"

type UserAssetsQuantityByTicker struct {
	Ticker   string `json:"ticker" db:"ticker"`
	Quantity int    `json:"quantity" db:"quantity"`
}

type UserAssetsValues struct {
	Ticker       string  `json:"ticker"`
	Quantity     int     `json:"quantity"`
	Value        float32 `json:"value"`
	ValuePerUnit float32 `json:"value_per_unit"`
}

type UserAssetsStocksInfo struct {
	CompanyStock      models.CompanyStock `json:"company_stock"`
	Quantity          int                 `json:"quantity"`
	BuyPrice          float32             `json:"buy_price"`
	CurrentPrice      float32             `json:"current_price"`
	TotalBuyPrice     float32             `json:"total_buy_price"`
	TotalCurrentPrice float32             `json:"total_current_price"`
	ProfitMargin      string              `json:"profit_margin"`
	ProfitPerUnit     float32             `json:"profit_per_unit"`
	TotalProfit       float32             `json:"total_profit"`
}

type UserAssetsProfits struct {
	UserAssets        []UserAssetsStocksInfo `json:"company_stock"`
	TotalSpent        float32                     `json:"total_spent"`
	TotalCurrentValue float32                     `json:"total_current_value"`
	TotalProfitMargin string                      `json:"total_profit_margin"`
	TotalProfit       float32                     `json:"total_profit_"`
}
