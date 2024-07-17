package models

type Service struct {
	Base
	ID          string  `json:"id,omitempty"`
	NameService string  `json:"nameservice"`
	ServiceDesc string  `json:"servicedesc"`
	Price       float64 `json:"price"`
	OwnerID     string  `json:"ownerid"`
}
