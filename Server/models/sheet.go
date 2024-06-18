package models

import "github.com/google/uuid"

type SheetStored struct {
	Base
	Name    string    `json:"sheetName"`
	OwnerID uuid.UUID `json:"ownerid"`
}
