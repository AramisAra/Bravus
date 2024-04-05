package models

import "time"

// Model for the Client table
type Client struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreateAt  time.Time
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Phone     uint   `json:"phone"`
	DogBreed  string `json:"dog_breed"`
	DogName   string `json:"dog_name"`
	DogAge    uint   `json:"dog_age"`
}
