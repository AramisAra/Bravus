package models

type Owner struct {
	Base
	Full_Name string `json:"full_name"`
	Phone     int    `json:"phone"`
	Email     string `json:"email"`
	Career    string `json:"career"`
}
