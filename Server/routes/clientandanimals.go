package routes

import (
	database "github.com/AramisAra/GroomingApp/database"
	models "github.com/AramisAra/GroomingApp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func isValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func CreateClientAndAnimal(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Insert client first
	result := database.Database.Db.Create(&client)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	// Get the newly created client ID
	clientID := client.ID

	var animal models.Animals

	animal.Client_id = clientID // Set foreign key to client ID
	if err := c.BodyParser(&animal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Insert animal with the foreign key set
	result = database.Database.Db.Create(&animal)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	Response := database.CreateJoinResult(client, animal)

	return c.Status(201).JSON(Response)
}

func ListClients(c *fiber.Ctx) error {
	clients := []models.Client{}

	database.Database.Db.Find(&clients)
	responseClients := []database.ClientSerializer{}

	for _, client := range clients {
		responseClients = append(responseClients, database.CreateClientResponse(client))
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

	if err := database.FindClient(parsedId, &client); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	// database.Database.Db.Where(&client, "id = ?", parsedId).First(&client)
	// database.Database.Db.Preload("Animals").Where(&client, "id = ?", parsedId).First(&client)
	responseClient := database.CreateClientResponse(client)

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
	if err := database.FindClient(parsedId, &client); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var updateClient database.UpdateClientInput

	if err := c.BodyParser(&updateClient); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	client.Full_Name = updateClient.FullName
	client.Email = updateClient.Email
	client.Phone = updateClient.Phone
	database.Database.Db.Save(&client)

	responseClient := database.CreateClientResponse(client)
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

	if err := database.FindClient(parsedId, &client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&client).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Client was deleted")
}

func CreateAnimal(c *fiber.Ctx) error {
	var animal models.Animals
	var client_id models.Client
	animal.Client_id = client_id.ID

	if err := c.BodyParser(&animal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&animal)
	responseAnimal := database.CreateAnimalResponse(animal)

	return c.Status(200).JSON(responseAnimal)
}

func ListAnimals(c *fiber.Ctx) error {
	animals := []models.Animals{}

	database.Database.Db.Find(&animals)
	responseAnimal := []database.AnimalSerializer{}

	for _, animal := range animals {
		responseAnimal = append(responseAnimal, database.CreateAnimalResponse(animal))
	}

	return c.Status(200).JSON(responseAnimal)
}

func GetAnimal(c *fiber.Ctx) error {
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

	animal := models.Animals{}

	if err := database.FindAnimal(parsedId, &animal); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseAnimal := database.CreateAnimalResponse(animal)

	return c.Status(200).JSON(responseAnimal)
}

func UpdateAnimal(c *fiber.Ctx) error {
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

	animal := models.Animals{}
	if err := database.FindAnimal(parsedId, &animal); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var updateAnimal database.UpdateAnimalInput

	if err := c.BodyParser(&updateAnimal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	animal.Animal_Name = updateAnimal.Animal_Name
	animal.Animal_Specie = updateAnimal.Animal_Specie
	animal.Animal_Age = updateAnimal.Animal_Age

	responseAnimal := database.CreateAnimalResponse(animal)
	return c.Status(200).JSON(responseAnimal)
}

func DeleteAnimal(c *fiber.Ctx) error {
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

	animal := models.Animals{}
	if err := database.FindAnimal(parsedId, &animal); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&animal).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Animal Was Deleted")
}
