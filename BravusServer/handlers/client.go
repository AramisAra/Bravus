package handlers

import (
	"github.com/AramisAra/BravusServer/database"
	"github.com/AramisAra/BravusServer/database/models"
	"github.com/gofiber/fiber/v2"
)

// CreateClient creates a new client in the database.
// It takes a fiber.Ctx object as a parameter and returns an error.
// The function parses the request body into a models.Client object,
// creates a new client in the database using the connectDb function,
// and returns the created client as a JSON response.
func CreateClient(c *fiber.Ctx) error {

	db := database.ConnectDb()

	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db.Create("Client", client)

	db.Close()

	return c.Status(200).JSON(client)
}

// ListClient handles the request to list all clients.
// It connects to the database, queries the "Client" table, and returns the response as JSON.
// ListClient handles the request to list all clients.
func ListClient(c *fiber.Ctx) error {

	db := database.ConnectDb()

	// query database
	response, err := db.Query("SELECT * FROM Client", map[string]interface{}{
		"tb": "Client",
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error})
	}

	db.Close()

	return c.Status(200).JSON(response)
}

func GetClient(c *fiber.Ctx) error {
	id := c.Query("id")

	db := database.ConnectDb()

	// Queries for the client
	response, err := db.Query("SELECT * From Client WHERE id = $id;", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db.Close()

	return c.Status(200).JSON(response)
}

// UpdateClient updates a client's information in the database based on the provided ID.
// It expects a JSON payload containing the updated client information in the request body.
// The updated client information includes the name, email, and phone number.
// It returns the updated client information as a JSON response.
func UpdateClient(c *fiber.Ctx) error {
	id := c.Query("id")

	db := database.ConnectDb()

	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	response, err := db.Update(id, map[string]interface{}{
		"name":  client.Name,
		"email": client.Email,
		"phone": client.Phone,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	db.Close()

	return c.Status(200).JSON(response)
}

func DeleteClient(c *fiber.Ctx) error {
	id := c.Query("id")

	db := database.ConnectDb()

	response, err := db.Query("DELETE FROM Client WHERE id = $id;", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db.Close()

	return c.Status(200).JSON(response)
}
