package models

import (
	"github.com/google/uuid"
)

// Model for animal table
type Animals struct {
	Base
	Animal_Name   string        `json:"animal_name"`
	Animal_Specie string        `json:"animal_specie"`
	Animal_Age    uint          `json:"animal_age"`
	Client_id     uuid.UUID     `json:"OwnerOfPet"`
	Appointments  []Appointment `gorm:"foreignKey:AnimalID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"appointments"`
}
