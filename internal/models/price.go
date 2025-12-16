package models

import "time"

type Price struct {
	Price     float64
	Provider  string
	UpdatedAt time.Time
}
