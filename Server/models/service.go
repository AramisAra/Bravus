package models

import "time"

// Model for the Service table
type Service struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreateAt    time.Time
	NameService string  `json:"name_service"`
	ServiceDesc string  `json:"service_desc"`
	Price       float64 `json:"price"`
	ServiceCode string  `json:"service_code"`
}
