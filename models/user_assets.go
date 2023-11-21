package models

import "time"

type UserAssets struct {
	ID             string     `json:"id" db:"id"`
	UserID         string     `json:"user_id" db:"user_id"`
	CompanyStockID string     `json:"company_stock_id" db:"company_stock_id"`
	Ticker         string     `json:"ticker" db:"ticker"`
	Quantity       int        `json:"quantity" db:"quantity"`
	PricePerUnit   float32    `json:"price_per_unit" db:"price_per_unit"`
	CreatedAt      *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
