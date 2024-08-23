package models

// Struct that holds the information of the Database
// This is an example of the structs info
//
//	{
//		"ID": "Client:304fks31sax"
//		"Name": "John Doe"
//		"Phone": "123-243-5677"
//		"Email": "nn@nn.com"
//		"Password": "HashPassword this will not show up in plaintext"
//	}
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
