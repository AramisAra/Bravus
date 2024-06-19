package handlers

import (
	"time"

	"github.com/AramisAra/BravusBackend/config"
	database "github.com/AramisAra/BravusBackend/database"
	dbmodels "github.com/AramisAra/BravusBackend/database/dbmodels"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func isValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

/*
	All client call return animal back with the client\
	the data of animal will be hide and shown as needed
*/
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

func LoginClient(c *fiber.Ctx) error {
	login := dbmodels.Login{}
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var client dbmodels.Client
	result := database.Database.Db.Find(&client, "Email = ?", login.Email)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(login.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid Password"})
	}

	claims := jwt.MapClaims{
		"ID":    client.ID,
		"email": client.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiry set to 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not login"})
	}

	// Return the token
	return c.JSON(dbmodels.LoginResponse{Token: t})

}

// Get Client related appointments
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

// Get Clients
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

// List all client
func ListClients(c *fiber.Ctx) error {
	clients := []dbmodels.Client{}

	database.Database.Db.Preload("Animals").Find(&clients)
	responseClients := []dbmodels.ClientSerializer{}

	for _, client := range clients {
		responseClients = append(responseClients, dbmodels.CreateClientResponse(client))
	}

	return c.Status(200).JSON(responseClients)
}

// Update client personal info
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

// Delete clients
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
