package handlers

import (
	"time"

	"github.com/AramisAra/BravusServer/config"
	"github.com/AramisAra/BravusServer/database"
	"github.com/AramisAra/BravusServer/database/models"
	"github.com/AramisAra/BravusServer/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
)

// CreateClient creates a new client in the database.
// It takes a fiber.Ctx object as a parameter and returns an error.
// The function parses the request body into a models.Client object,
// creates a new client in the database using the connectDb function,
// and returns the created client as a JSON response.
func RegisterClient(c *fiber.Ctx) error {

	db := database.ConnectDb()
	defer db.Close()
	var client models.Client
	createdAt := time.Now().Format(time.RFC3339)

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(client.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Failed to crypt"})
	}
	client.Password = hashedPassword

	// Construct and execute the query with placeholders and a map for the values
	query := "CREATE Client CONTENT {name: $name, email: $email, password: $password, phone: $phone, createdAt: $createdAt, updateAt: $updateAt, animals: $animals, appointments: $appointments}"
	values := map[string]interface{}{
		"name":         client.Name,
		"email":        client.Email,
		"password":     client.Password,
		"phone":        client.Phone,
		"createdAt":    createdAt,
		"updateAt":     createdAt,
		"animals":      client.Animals,
		"appointments": client.Appointments,
	}

	_, err = db.Query(query, values)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(client)
}

func LoginClient(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db := database.ConnectDb()
	defer db.Close() // Ensure the database connection is closed after the query

	// Query for the client by email
	query := `SELECT * FROM Client WHERE email = $email`
	params := map[string]interface{}{
		"email": input.Email,
	}

	response, err := db.Query(query, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Assuming response is a JSON array, unmarshal into a slice of maps
	var results []map[string]interface{}
	if err := surrealdb.Unmarshal(response, &results); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if the results slice is empty
	if len(results) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	// Extract the password from the nested structure
	data, ok := results[0]["result"].([]interface{})
	if !ok || len(data) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	clientData, ok := data[0].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "password retrieval error"})
	}

	storedPassword, ok := clientData["password"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "password retrieval error"})
	}
	if err := utils.ComparePasswords(storedPassword, input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	clientid := clientData["id"].(string)
	clientemail := clientData["email"].(string)
	claims := jwt.MapClaims{
		"ID":    clientid,
		"email": clientemail,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not login"})
	}

	// Return the token
	return c.JSON(models.LoginResponse{ID: clientid, Token: t})
}

// ListClient handles the request to list all clients.
// It connects to the database, queries the "Client" table, and returns the response as JSON.
// ListClient handles the request to list all clients.
func ListClient(c *fiber.Ctx) error {
	db := database.ConnectDb()
	defer db.Close()
	// query database
	response, err := db.Query("SELECT * FROM Client", map[string]interface{}{
		"tb": "Client",
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error})
	}

	return c.Status(200).JSON(response)
}

func GetClient(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No ID"})
	} else {
		checker := utils.IsValidClientString(id)
		if !checker {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not a valid id"})
		}
	}
	db := database.ConnectDb()
	defer db.Close()

	// Queries for the client
	response, err := db.Query("SELECT * From Client WHERE id = $id;", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(response)
}

// UpdateClient updates a client's information in the database based on the provided ID.
// It expects a JSON payload containing the updated client information in the request body.
// The updated client information includes the name, email, and phone number.
// It returns the updated client information as a JSON response.
func UpdateClient(c *fiber.Ctx) error {
	id := c.Query("id")

	db := database.ConnectDb()
	defer db.Close()

	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}

	updateAt := time.Now().Format(time.RFC3339)

	query := "UPDATE Client:$id CONTENT {name: $name, email: $email, phone: $phone, createdAt: $createdAt, updateAt: $updateAt, animals: $animals, appointments: $appointments}"
	values := map[string]interface{}{
		"id":           id,
		"name":         client.Name,
		"email":        client.Email,
		"phone":        client.Phone,
		"updateAt":     updateAt,
		"animals":      client.Animals,
		"appointments": client.Appointments,
	}

	db.Query(query, values)

	return c.Status(200).JSON(client)
}

func DeleteClient(c *fiber.Ctx) error {
	id := c.Query("id")

	db := database.ConnectDb()
	defer db.Close()

	response, err := db.Query("DELETE FROM Client WHERE id = $id;", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(response)
}
