package dbmodels

import (
	"github.com/google/uuid"
)

type joinResultClient struct {
	Client Client `josn:"client"`
	Animal Animal `json:"animal"`
}

type ClientSerializer struct {
	ID           uuid.UUID     `json:"id"`
	FullName     string        `json:"full_name"`
	Email        string        `json:"email"`
	Phone        uint          `json:"phone"`
	Animals      []Animal      `json:"animals"`
	Appointments []Appointment `json:"appointments"`
}

type UpdateClientInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    uint   `json:"phone"`
}

type ServiceSerializer struct {
	ID          uuid.UUID `json:"id"`
	ServiceName string    `json:"service_name"`
	ServiceDesc string    `json:"service_desc"`
	ServiceCode string    `json:"service_code"`
}

type UpdateServiceInput struct {
	ServiceName string  `json:"service_name"`
	ServiceDesc string  `json:"service_desc"`
	Price       float64 `json:"price"`
	ServiceCode string  `json:"service_code"`
}

type OwnerSerializer struct {
	ID           uuid.UUID     `json:"id"`
	Full_Name    string        `json:"full_name"`
	Phone        uint          `json:"phone"`
	Email        string        `json:"email"`
	Career       string        `json:"career"`
	Appointments []Appointment `json:"appointments"`
}

type UpdateOwnerInput struct {
	Full_Name string `json:"full_name"`
	Phone     uint   `json:"phone"`
	Email     string `json:"email"`
	Career    string `json:"career"`
}

type AnimalSerializer struct {
	ID            uuid.UUID `json:"id"`
	Animal_Name   string    `json:"animal_name"`
	Animal_Specie string    `json:"animal_specie"`
	Animal_Age    uint      `json:"animal_age"`
	Client        uuid.UUID `json:"OwnerOfAnimal"`
}

type UpdateAnimalInput struct {
	Animal_Name   string `json:"animal_name"`
	Animal_Specie string `json:"animal_specie"`
	Animal_Age    uint   `json:"animal_age"`
}

type AppointmentSerializer struct {
	ClientID uuid.UUID `json:"client"`
	OwnerID  uuid.UUID `json:"owner"`
	AnimalID uuid.UUID `json:"animal"`
	Date     string    `json:"date"`
	Time     string    `json:"time"`
}

type AppointmentUpdater struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
