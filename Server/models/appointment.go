package models

import (
	"time"

	"github.com/google/uuid"
)

// Models for the appointment table
type Appointment struct {
	Base
	Date     time.Time `json:"date" gorm:"DATE"`
	Time     time.Time `json:"time" gorm:"TIME"`
	OwnerID  uuid.UUID `json:"ownerid"`
	ClientID uuid.UUID `json:"clientid"`
	AnimalID uuid.UUID `json:"animalid"`
	Services []Service `gorm:"foreignKey:AppointmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"serviceid"`
}
