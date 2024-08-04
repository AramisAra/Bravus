package models

type Client struct {
	Base
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name"`
	Phone        string        `json:"phone"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Animals      []Animal      `json:"animals"`
	Appointments []Appointment `json:"appointment"`
}
