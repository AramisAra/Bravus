package utils

import "github.com/AramisAra/BravusServer/database/models"

type ClientSerializer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func CreateClientResponse(client models.Client) ClientSerializer {
	return ClientSerializer{ID: client.ID, Name: client.Name, Email: client.Email,
		Phone: client.Phone}
}
