package utils

import (
	"errors"

	"github.com/AramisAra/GroomingApp/database"
	"github.com/AramisAra/GroomingApp/models"
	"github.com/google/uuid"
)

// Client's ulits code
type ClientSerializer struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Phone    uint      `json:"phone"`
}

type UpdateClientInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    uint   `json:"phone"`
}

func CreateClientResponse(client models.Client) ClientSerializer {
	return ClientSerializer{ID: client.ID, FullName: client.Full_Name, Email: client.Email, Phone: client.Phone}
}

func FindClient(id uuid.UUID, client *models.Client) error {
	database.Database.Db.Find(&client, "id = ?", id)

	if client.ID == uuid.Nil {
		return errors.New("couldn't find client")
	}
	return nil
}

// Service's ulits code

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

func CreateServiceResponse(service models.Service) ServiceSerializer {
	return ServiceSerializer{ID: service.ID, ServiceName: service.NameService, ServiceDesc: service.ServiceDesc, ServiceCode: service.ServiceCode}
}

func FindService(id uuid.UUID, service *models.Service) error {
	database.Database.Db.Find(&service, "id = ?", id)
	if service.ID == uuid.Nil {
		return errors.New("couldn't find service")
	}
	return nil
}

// Owner's ulits code

type OwnerSerializer struct {
	ID        uuid.UUID `json:"id"`
	Full_Name string    `json:"full_name"`
	Phone     uint      `json:"phone"`
	Email     string    `json:"email"`
	Career    string    `json:"career"`
}

type UpdateOwnerInput struct {
	Full_Name string `json:"full_name"`
	Phone     uint   `json:"phone"`
	Email     string `json:"email"`
	Career    string `json:"career"`
}

func CreateOwnerResponse(owner models.Owner) OwnerSerializer {
	return OwnerSerializer{ID: owner.ID, Full_Name: owner.Full_Name, Phone: owner.Phone, Email: owner.Email, Career: owner.Career}
}

func FindOwner(id uuid.UUID, owner *models.Owner) error {
	database.Database.Db.Find(&owner, "id = ?", id)
	if owner.ID == uuid.Nil {
		return errors.New("couldn't find owner")
	}
	return nil
}

// Animal's ulits code

type AnimalSerializer struct {
	ID            uuid.UUID `json:"id"`
	Animal_Name   string    `json:"animal_name"`
	Animal_Specie string    `json:"animal_specie"`
	Animal_Age    string    `json:"animal_age"`
}

type UpdateAnimalInput struct {
	Animal_Name   string `json:"animal_name"`
	Animal_Specie string `json:"animal_specie"`
	Animal_Age    string `json:"animal_age"`
}

func CreateAnimalResponse(animal models.Animals) AnimalSerializer {
	return AnimalSerializer{ID: animal.ID, Animal_Name: animal.Animal_Name, Animal_Specie: animal.Animal_Specie, Animal_Age: animal.Animal_Age}
}

func FindAnimal(id uuid.UUID, animal *models.Animals) error {
	database.Database.Db.Find(&animal, "id = ?", id)
	if animal.ID == uuid.Nil {
		return errors.New("couldn't find animal")
	}
	return nil
}
