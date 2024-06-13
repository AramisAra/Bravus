package models

// Model for the Client table
type Client struct {
	Base
	Full_Name   string    `json:"full_name"`
	Email       string    `json:"email"`
	Phone       uint      `json:"phone"`
	AnimalRefer []Animals `gorm:"foreignKey:Client_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Animals     Animals   `json:"animals"`
}
