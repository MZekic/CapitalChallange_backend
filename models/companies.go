package models

import "time"

type Companies struct {
	ID              string     `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	Ticker          string     `json:"ticker" db:"ticker"`
	Market          string     `json:"market" db:"market"`
	PrimaryExchange string     `json:"primary_exchange" db:"primary_exchange"`
	Type            string     `json:"type" db:"type"`
	CurrencyName    string     `json:"currency_name" db:"currency_name"`
	Locale          string     `json:"locale" db:"locale"`
	HomepageURL     string     `json:"homepage_url" db:"homepage_url"`
	Description     string     `json:"description" db:"description"`
	TotalEmployees  float32    `json:"total_employees" db:"total_employees"`
	LogoURL         string     `json:"logo_url" db:"logo_url"`
	CreatedAt       *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
