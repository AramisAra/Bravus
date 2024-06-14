package models

// Model for the Client table
type Client struct {
	Base
	Full_Name    string        `json:"full_name"`
	Email        string        `json:"email"`
	Phone        uint          `json:"phone"`
	Appointments []Appointment `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"appointments"`
	Animals      []Animals     `gorm:"foreignKey:Client_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"animals"`
}
