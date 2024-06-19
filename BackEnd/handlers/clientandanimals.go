package handlers

import (
	database "github.com/AramisAra/BravusBackend/database"
	dbmodels "github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func isValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// handles Register and Client creation
func RegisterClient(c *fiber.Ctx) error {
	registration := new(dbmodels.RegisterRequestClient)
	if err := c.BodyParser(registration); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if the email already exists
	var existingUser dbmodels.Client
	if err := database.Database.Db.Where("email = ?", registration.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already registered"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not hash password"})
	}

	// Create the user
	client := dbmodels.Client{
		Full_Name: registration.Name,
		Email:     registration.Email,
		Phone:     registration.Phone,
		Password:  string(hashedPassword),
	}

	// Save the user in the database
	result := database.Database.Db.Create(&client)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	return c.Status(fiber.StatusCreated).JSON(client)
}

// For Creating A Second Animals
func CreateAnimal(c *fiber.Ctx) error {
	// Creates Client first
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	client := dbmodels.Client{}

	database.Database.Db.Find(&client, "id = ?", id)

	ClientID := client.ID

	var animal dbmodels.Animal

	animal.Client_id = ClientID
	if err := c.BodyParser(&animal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.Database.Db.Create(&animal)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	Response := dbmodels.CreateAnimalResponse(animal)

	return c.Status(201).JSON(Response)
}

// Get Clients with Animals
func GetClient(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	client := dbmodels.Client{}

	database.Database.Db.Preload("Animals").Find(&client, "id = ?", id)

	response := dbmodels.CreateClientResponse(client)

	return c.Status(200).JSON(response)
}

func ListClients(c *fiber.Ctx) error {
	clients := []dbmodels.Client{}

	database.Database.Db.Preload("Animals").Find(&clients)
	responseClients := []dbmodels.ClientSerializer{}

	for _, client := range clients {
		responseClients = append(responseClients, dbmodels.CreateClientResponse(client))
	}

	return c.Status(200).JSON(responseClients)
}

func UpdateClient(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	client := dbmodels.Client{}

	database.Database.Db.Joins("Animals").Find(&client, "id = ?", id)

	var updateClient dbmodels.UpdateClientInput

	if err := c.BodyParser(&updateClient); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	client.Full_Name = updateClient.FullName
	client.Email = updateClient.Email
	client.Phone = updateClient.Phone

	database.Database.Db.Omit("Animals").Save(&client)

	responseClient := dbmodels.CreateClientResponse(client)
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

	client := dbmodels.Client{}

	database.Database.Db.Find(&client)
	if err := database.Database.Db.Delete(&client).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Client was deleted")
}

func DeleteAnimal(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	animal := dbmodels.Animal{}

	database.Database.Db.Find(&animal, "id = ?", id)
	if err := database.Database.Db.Delete(&animal).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Animal was deleted")
}

/*
func ListAnimals(c *fiber.Ctx) error {
	animals := []dbmodels.Animal{}

	database.Database.Db.Find(&animals)
	responseAnimal := []dbmodels.AnimalSerializer{}

	for _, animal := range animals {
		responseAnimal = append(responseAnimal, dbmodels.CreateAnimalResponse(animal))
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

	animal := dbmodels.Animal{}

	database.Database.Db.Find(&animal)
	responseAnimal := dbmodels.CreateAnimalResponse(animal)

	return c.Status(200).JSON(responseAnimal)
}
*/

func UpdateAnimal(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	animal := dbmodels.Animal{}

	database.Database.Db.Find(&animal, "id = ?", id)

	var updateAnimal dbmodels.UpdateAnimalInput

	if err := c.BodyParser(&updateAnimal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	animal.Animal_Name = updateAnimal.Animal_Name
	animal.Animal_Specie = updateAnimal.Animal_Specie
	animal.Animal_Age = updateAnimal.Animal_Age

	database.Database.Db.Save(&animal)

	responseAnimal := dbmodels.CreateAnimalResponse(animal)
	return c.Status(200).JSON(responseAnimal)
}

func GetAppointment(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	var client dbmodels.Client

	database.Database.Db.Preload("Appointments").Preload("Animals").Find(&client)

	response := dbmodels.CreateClientResponse(client)

	return c.Status(200).JSON(response)
}
