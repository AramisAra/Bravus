package dbmodels

// Model for the Client table
type Client struct {
	Base
	Full_Name    string        `json:"full_name"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Phone        uint          `json:"phone"`
	Appointments []Appointment `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"appointments"`
	Animals      []Animal      `gorm:"foreignKey:Client_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"animals"`
}
