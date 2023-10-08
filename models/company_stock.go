package models

import "time"

type CompanyStock struct {
	ID                         string     `json:"id" db:"id"`
	Ticker                     string     `json:"ticker" db:"ticker"`
	ClosePrice                 float32    `json:"close_price" db:"close_price"`
	HighestPrice               float32    `json:"highest_price" db:"highest_price"`
	LowestPrice                float32    `json:"lowest_price" db:"lowest_price"`
	NumberOfTransactions       int        `json:"number_of_transactions" db:"number_of_transactions"`
	OpenPrice                  float32    `json:"open_price" db:"open_price"`
	OTC                        bool       `json:"otc" db:"otc"`
	TradingVolume              int        `json:"trading_volume" db:"trading_volume"`
	VolumeWeightedAveragePrice float32    `json:"volume_weighted_average_price" db:"volume_weighted_average_price"`
	CreatedAt                  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt                  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
