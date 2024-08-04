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
		"createdAt":    createdAt,
		"updateAt":     createdAt,
		"name":         client.Name,
		"phone":        client.Phone,
		"email":        client.Email,
		"password":     client.Password,
		"animals":      client.Animals,
		"appointments": client.Appointments,
	}

	_, err = db.Query(query, values)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(client)
}

// LoginClient handles the login request for a client.
// It expects a JSON payload with the following structure:
//
//	{
//	  "email": "user@example.com",
//	  "password": "password123"
//	}
//
// It retrieves the client from the database based on the provided email,
// verifies the password, and generates a JWT token if the login is successful.
// The token is returned in the response along with the client ID.
// If the email or password is invalid, it returns an error response.
func LoginClient(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db := database.ConnectDb()
	defer db.Close()

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

// UpdateClient updates the client information in the database based on the provided ID.
// It expects a query parameter 'id' to be passed in the request URL.
// If the 'id' is empty or not a valid ID, it returns a JSON response with an error message.
// It fetches the existing client data from the database using the provided ID.
// If the client is not found, it returns a JSON response with an error message.
// It parses the request body into a client struct and updates the corresponding fields in the existing client data.
// It updates the 'updateAt' field with the current timestamp.
// It builds an update query and executes it to update the client information in the database.
// Finally, it returns a JSON response with the updated client data.
func UpdateClient(c *fiber.Ctx) error {
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

	// Fetch the existing client data from the database
	query := `SELECT * FROM Client WHERE id = $id`
	params := map[string]interface{}{
		"id": id,
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
	if len(results) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Client not found"})
	}

	// Extract the existing client data from the nested structure
	existingData, ok := results[0]["result"].([]interface{})
	if !ok || len(existingData) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid client data format"})
	}
	existingClientData, ok := existingData[0].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid client data format"})
	}

	// Parse the request body into the client struct
	var client models.Client
	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}

	// Only update the fields that are provided in the request body
	if client.Name != "" {
		existingClientData["name"] = client.Name
	}
	if client.Email != "" {
		existingClientData["email"] = client.Email
	}
	if client.Phone != "" {
		existingClientData["phone"] = client.Phone
	}
	if client.Animals != nil {
		existingClientData["animals"] = client.Animals
	}
	if client.Appointments != nil {
		existingClientData["appointments"] = client.Appointments
	}

	// Update the updateAt field
	updateAt := time.Now().Format(time.RFC3339)

	// Build the update query
	updateQuery := "UPDATE $id CONTENT {name: $name, email: $email, phone: $phone, password: $password, createdAt: $createdAt, updateAt: $updateAt, animals: $animals, appointments: $appointments}"
	values := map[string]interface{}{
		"id":           id,
		"createdAt":    existingClientData["createdAt"],
		"updateAt":     updateAt,
		"name":         existingClientData["name"],
		"phone":        existingClientData["phone"],
		"email":        existingClientData["email"],
		"password":     existingClientData["password"],
		"animals":      existingClientData["animals"],
		"appointments": existingClientData["appointments"],
	}

	// Execute the update query
	if _, err := db.Query(updateQuery, values); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(existingClientData)
}

func DeleteClient(c *fiber.Ctx) error {
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

	response, err := db.Query("DELETE FROM Client WHERE id = $id;", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(response)
}
