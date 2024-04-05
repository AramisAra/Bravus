package models

import "time"

// Model for the Client table
type Client struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreateAt  time.Time
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
}
