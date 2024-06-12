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

func FindClient(id int, client *models.Client) error {
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

func FindService(id int, service *models.Service) error {
	database.Database.Db.Find(&service, "id = ?", id)
	if service.ID == uuid.Nil {
		return errors.New("couldn't find service")
	}
	return nil
}
