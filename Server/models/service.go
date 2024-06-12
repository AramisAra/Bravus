package models

import "github.com/google/uuid"

// Model for the Service table
type Service struct {
	Base
	AppointmentRefer uuid.UUID
	NameService      string  `json:"name_service"`
	ServiceDesc      string  `json:"service_desc"`
	Price            float64 `json:"price"`
	ServiceCode      string  `json:"service_code"`
}
