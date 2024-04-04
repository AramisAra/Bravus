package routes

import (
	database "github.com/AramisAra/Grooming_App/Database"
	models "github.com/AramisAra/Grooming_App/Database/Models"
	"github.com/gofiber/fiber/v2"
)

type ClientSerializer struct {
	// This is not the model, This is for serialtion
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseClient(clientModel models.Client) ClientSerializer {
	return ClientSerializer{ID: clientModel.ID, FirstName: clientModel.FirstName, LastName: clientModel.LastName}
}

func CreateClient(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&client)
	responseClient := CreateResponseClient(client)

	return c.Status(200).JSON(responseClient)
}

func GetClients(c *fiber.Ctx) error {
	clients := []models.Client{}

	database.Database.Db.Find(&clients)
	responseClients := []ClientSerializer{}
	for _, client := range clients {
		responseClient := CreateResponseClient(client)
		responseClients = append(responseClients, responseClient)
	}

	return c.Status(200).JSON(responseClients)
}
