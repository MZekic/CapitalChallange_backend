package models

import "time"

type UserTransactions struct {
	ID             string     `json:"id" db:"id"`
	UserID         string     `json:"user_id" db:"user_id"`
	CompanyStockID string     `json:"company_stock_id" db:"company_stock_id"`
	BuyPrice       float32    `json:"buy_price" db:"buy_price"`
	SellPrice      float32    `json:"sell_price" db:"sell_price"`
	Quantity       int        `json:"quantity" db:"quantity"`
	CreatedAt      *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type UserTransactionUnit struct {
	Ticker      	  string  `json:"ticker" db:"ticker"`
	CompanyStockID string     `json:"company_stock_id" db:"company_stock_id"`
	BuyPrice       float32    `json:"buy_price" db:"buy_price"`
	SellPrice      float32    `json:"sell_price" db:"sell_price"`
	Quantity       int        `json:"quantity" db:"quantity"`
	CreatedAt      *time.Time `json:"created_at,omitempty" db:"created_at"`
}
