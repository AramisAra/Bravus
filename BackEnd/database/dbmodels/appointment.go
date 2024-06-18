package dbmodels

import (
	"github.com/google/uuid"
)

// Models for the appointment table
type Appointment struct {
	Base
	Date     string    `json:"date" gorm:"DATE"`
	Time     string    `json:"time" gorm:"TIME"`
	OwnerID  uuid.UUID `json:"ownerid"`
	ClientID uuid.UUID `json:"clientid"`
	AnimalID uuid.UUID `json:"animalid"`
}
