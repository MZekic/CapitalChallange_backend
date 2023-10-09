package models

import "time"

type UserBalance struct {
	ID              string     `json:"id" db:"id"`
	UserID          string     `json:"user_id" db:"user_id"`
	StartingBalance float32    `json:"starting_balance" db:"starting_balance"`
	CurrentBalance  float32    `json:"current_balance" db:"current_balance"`
	GameNumber      int        `json:"game_number" db:"game_number"`
	CreatedAt       *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
