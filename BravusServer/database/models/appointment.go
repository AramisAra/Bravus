package models

// Models for the appointment table
type Appointment struct {
	Base
	ID       string `json:"id,omitempty"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	OwnerID  string `json:"ownerid"`
	ClientID string `json:"clientid"`
}
