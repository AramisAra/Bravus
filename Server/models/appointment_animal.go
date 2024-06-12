package models

import (
	"time"

	"github.com/google/uuid"
)

type Appointment_Animal struct {
	Base
	Date              time.Time `json:"date" gorm:"DATE"`
	Time              time.Time `json:"time" gorm:"TIME"`
	ProfessionalRefer uuid.UUID `json:"professional_id"`
	Professional      Owner     `gorm:"type:uuid;foreignKey:ProfessionalRefer"`
	ClientRefer       uuid.UUID `json:"client_id"`
	Client            Client    `gorm:"type:uuid;foreignKey:CleintRefer"`
	AnimalRefer       uuid.UUID `json:"animal_id"`
	Animal            Animals   `gorm:"type:uuid;foreignKey:AnimalRefer"`
	ServiceRefer      uuid.UUID `json:"service_id"`
	Service           Service   `gorm:"type:uuid;foreignKey:ServiceRefer"`
}
