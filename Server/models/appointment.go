package models

import "time"

// Models for the appointment table
type Appointment struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreateAt     time.Time
	ClientRefer  int     `json:"client_id"`
	Client       Client  `gorm:"foreignKey:ClientRefer"`
	ServiceRefer int     `json:"service_id"`
	Service      Service `gorm:"foreignKey:ServiceRefer"`
}
