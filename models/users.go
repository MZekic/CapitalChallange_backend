package models

import "time"

type Users struct {
	ID                string     `json:"id" db:"id"`
	Username          string     `json:"username" db:"username"`
	Password          string     `json:"password" db:"password"`
	Email             string     `json:"email" db:"email"`
	CurrentGameNumber int        `json:"current_game_number" db:"current_game_number"`
	CreatedAt         *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
