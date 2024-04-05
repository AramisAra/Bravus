package utils

import (
	"errors"

	"github.com/AramisAra/GroomingApp/database"
	"github.com/AramisAra/GroomingApp/models"
)

type ClientSerializer struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    uint   `json:"phone"`
	DogBreed string `json:"dog_breed"`
	DogName  string `json:"dog_name"`
	DogAge   uint   `json:"dog_age"`
}

func CreateClientResponse(client models.Client) ClientSerializer {
	return ClientSerializer{ID: client.ID, FullName: client.Full_Name, Email: client.Email, Phone: client.Phone, DogBreed: client.DogBreed, DogName: client.DogName, DogAge: client.DogAge}
}

func FindClient(id int, client *models.Client) error {
	database.Database.Db.Find(&client, "id = ?", id)

	if err != nil {
		return errors.New("Client not found")
	}
	return nil
}
