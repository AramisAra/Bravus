package main

import (
	database "github.com/AramisAra/Grooming_App/Database"
	routes "github.com/AramisAra/Grooming_App/Database/routes"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func setupRoutes(app *fiber.App) {
	// The following are endpoints
	// HealthCheck Endpoints
	app.Get("/HealthCheck", HealthCheck)
	// Client Endpoints
	app.Post("/Tests/CreateClient", routes.CreateClient)
	app.Get("/Tests/ListClients", routes.GetClients)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":8000")
}
