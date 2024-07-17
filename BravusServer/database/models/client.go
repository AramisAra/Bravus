package models

type Client struct {
	ID string `json:"id,omitempty"`
	Base
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
