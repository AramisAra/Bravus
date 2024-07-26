package models

type Owner struct {
	Base
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name"`
	Phone        string        `json:"phone"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Career       string        `json:"career"`
	Services     []Service     `json:"services"`
	Appointments []Appointment `json:"appointments"`
}
