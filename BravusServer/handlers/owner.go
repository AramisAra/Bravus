package handlers

import (
	"time"

	"github.com/AramisAra/BravusBackend/config"
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// RegisterOwner registers a new owner in the system.
// It takes a fiber.Ctx object as a parameter and returns an error.
// The function parses the request body to retrieve the registration details.
// It checks if the email already exists in the database for either an owner or a client.
// If the email already exists, it returns an error response.
// The function then hashes the password using bcrypt.
// It creates a new owner object with the registration details.
// Finally, it saves the owner in the database and returns the created owner as a JSON response.
func RegisterOwner(c *fiber.Ctx) error {
	registration := new(dbmodels.RegisterRequestOwner)
	if err := c.BodyParser(registration); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if the email already exists
	var existingOwner dbmodels.Owner
	var existingClient dbmodels.Client
	if err := database.Database.Db.Where("email = ?", registration.Email).First(&existingOwner).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already registered"})
	}
	if err := database.Database.Db.Where("email = ?", registration.Email).First(&existingClient).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already registered"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not hash password"})
	}

	// Create the user
	owner := dbmodels.Owner{
		Full_Name: registration.Name,
		Email:     registration.Email,
		Phone:     registration.Phone,
		Password:  string(hashedPassword),
		Career:    registration.Career,
	}

	// Save the user in the database
	result := database.Database.Db.Create(&owner)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	return c.Status(fiber.StatusCreated).JSON(owner)
}

// LoginOwner handles the login request for an owner.
// It receives a fiber.Ctx object representing the HTTP request context.
// The function parses the request body to retrieve the login credentials.
// It checks if the email exists in the database and verifies the password.
// If the credentials are valid, it generates a JWT token with the owner's ID and email.
// The token is signed using the secret key from the configuration.
// Finally, it returns the owner's ID and the generated token as a JSON response.
func LoginOwner(c *fiber.Ctx) error {
	login := dbmodels.Login{}
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Checkes if the email is in the database
	var owner dbmodels.Owner
	result := database.Database.Db.Find(&owner, "Email = ?", login.Email)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	// Checkes hash password
	if err := bcrypt.CompareHashAndPassword([]byte(owner.Password), []byte(login.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid Password"})
	}

	// Creates the Claims
	claims := jwt.MapClaims{
		"ID":    owner.ID,
		"email": owner.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiry set to 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not login"})
	}

	// Return the token
	return c.JSON(dbmodels.LoginResponse{ID: owner.ID, Token: t})
}

// ListOwners retrieves a list of owners from the database and returns them as a JSON response.
func ListOwners(c *fiber.Ctx) error {
	owners := []dbmodels.Owner{}

	database.Database.Db.Preload("Services").Find(&owners)
	responseOwner := []dbmodels.OwnerSerializer{}

	for _, owner := range owners {
		responseOwner = append(responseOwner, dbmodels.CreateOwnerResponse(owner))
	}

	return c.Status(200).JSON(responseOwner)
}

func GetOwner(c *fiber.Ctx) error {
	id := c.Query("uuid")
	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}

	owner := dbmodels.Owner{}

	database.Database.Db.Preload("Appointment").Find(&owner, "id = ?", id)

	response := dbmodels.CreateOwnerResponse(owner)

	return c.Status(200).JSON(response)
}

func GetAppointmentOwner(c *fiber.Ctx) error {
	id := c.Query("uuid")
	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	var appointments []dbmodels.Appointment
	database.Database.Db.Find(&appointments, "owner_id = ?", id)

	responseAppointment := []dbmodels.AppointmentSerializer{}
	for _, appointment := range appointments {
		responseAppointment = append(responseAppointment, dbmodels.CreateAppointmentResponse(appointment))
	}
	return c.Status(200).JSON(responseAppointment)
}

// UpdateOwner updates the information of an owner in the database.
// It takes a fiber.Ctx object as a parameter and returns an error.
// The function first checks if the provided UUID is valid.
// If the UUID is invalid, it returns a fiber.StatusBadRequest response with the message "Invalid UUID".
// It then retrieves the owner from the database based on the provided UUID.
// The function parses the request body into a dbmodels.UpdateOwnerInput object.
// If there is an error while parsing the request body, it returns a fiber.StatusBadRequest response with the error message.
// The function updates the owner's information with the values from the parsed request body.
// It saves the updated owner to the database, omitting the "Animals" association.
// Finally, it creates a dbmodels.CreateOwnerResponse object from the updated owner and returns a fiber.StatusOK response with the JSON representation of the response object.
func UpdateOwner(c *fiber.Ctx) error {
	id := c.Query("uuid")
	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	owner := dbmodels.Owner{}

	database.Database.Db.Joins("Services").Find(&owner, "id = ?", id)

	var updateowner dbmodels.UpdateOwnerInput

	if err := c.BodyParser(&updateowner); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	owner.Full_Name = updateowner.Full_Name
	owner.Email = updateowner.Email
	owner.Phone = updateowner.Phone
	owner.Career = updateowner.Career

	database.Database.Db.Omit("Animals").Save(&owner)

	responseowner := dbmodels.CreateOwnerResponse(owner)
	return c.Status(200).JSON(responseowner)
}

// DeleteOwner deletes an owner from the database based on the provided UUID.
// It returns an error if the UUID is invalid or if there was an issue deleting the owner.
// If the owner is successfully deleted, it returns a JSON response with a status of 200.
func DeleteOwner(c *fiber.Ctx) error {
	id := c.Query("uuid")
	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	owner := dbmodels.Owner{}

	database.Database.Db.Find(&owner)
	if err := database.Database.Db.Delete(&owner).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("owner was deleted")
}
