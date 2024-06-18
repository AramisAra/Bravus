package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base ID model
type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewV6()
	if err != nil {
		return err
	}
	b.ID = uuid
	return nil
}
