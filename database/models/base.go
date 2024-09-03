package models

import "time"

// This struct hold the base info on all the tables
type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
