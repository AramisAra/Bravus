package routes

import (
	database "github.com/AramisAra/GroomingApp/database"
	models "github.com/AramisAra/GroomingApp/models"
	"github.com/AramisAra/GroomingApp/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func isValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func CreateClient(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&client)
	responseClient := utils.CreateClientResponse(client)

	return c.Status(200).JSON(responseClient)
}

func ListClients(c *fiber.Ctx) error {
	clients := []models.Client{}

	database.Database.Db.Find(&clients)
	responseClients := []utils.ClientSerializer{}

	for _, client := range clients {
		responseClients = append(responseClients, utils.CreateClientResponse(client))
	}

	return c.Status(200).JSON(responseClients)
}

func GetClient(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	client := models.Client{}

	if err := utils.FindClient(parsedId, &client); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseClient := utils.CreateClientResponse(client)

	return c.Status(200).JSON(responseClient)
}

func UpdateClient(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	client := models.Client{}
	if err := utils.FindClient(parsedId, &client); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var updateClient utils.UpdateClientInput

	if err := c.BodyParser(&updateClient); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	client.Full_Name = updateClient.FullName
	client.Email = updateClient.Email
	client.Phone = updateClient.Phone
	database.Database.Db.Save(&client)

	responseClient := utils.CreateClientResponse(client)
	return c.Status(200).JSON(responseClient)
}

func DeleteClient(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	client := models.Client{}

	if err := utils.FindClient(parsedId, &client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&client).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Client was deleted")
}
