package entity

import (
	"time"
)

type GeoObject struct {
	ID        string
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
