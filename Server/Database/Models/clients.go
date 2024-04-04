package models

import "time"

// Model for the Client table
type Client struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreateAt  time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
