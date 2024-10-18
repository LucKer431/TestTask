package main

import (
	"time"
)

type Building struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	YearBuilt int       `json:"year_built"`
	Floors    int       `json:"floors"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
