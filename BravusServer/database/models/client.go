package models

type Client struct {
	Base
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Animals      []Animal      `json:"animals"`
	Appointments []Appointment `json:"appointment"`
}
