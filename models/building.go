package models

import (
	"time"
)

// @Description Building structure
type Building struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	City       string    `json:"city"`
	YearBuilt  int       `json:"year_built" swaggertype:"integer"`
	Floors     int       `json:"floors" swaggertype:"integer"`
	CreateTime time.Time `json:"created_at" swaggertype:"string"`
}