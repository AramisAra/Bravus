package models

import "time"

// Model for the Service table
type Service struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreateAt    time.Time
	NameService string `json:"name_service"`
	ServiceCode string `json:"service_code"`
}
