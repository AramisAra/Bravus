package models

// Models for Owner table
type Owner struct {
	Base
	Full_Name string `json:"full_name"`
	Phone     uint   `json:"phone"`
	Email     string `json:"email"`
	Career    string `json:"career"`
}
