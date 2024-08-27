package models

// This struct defines the animal table
type Animal struct {
	Base
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Species string `json:"species"`
	OwnerID string `json:"owner_id"`
}
