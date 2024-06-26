package dbmodels

// Models for Owner table
type Owner struct {
	Base
	Full_Name    string        `json:"full_name"`
	Phone        string        `json:"phone"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Career       string        `json:"career"`
	Services     []Service     `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"services"`
	Appointments []Appointment `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"appointments"`
}
