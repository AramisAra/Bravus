package main

import (
	database "github.com/AramisAra/GroomingApp/database"
	routes "github.com/AramisAra/GroomingApp/routes"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func setupRoutes(app *fiber.App) {
	// HealthCheck
	app.Get("/health", HealthCheck)
	// Client's dataRoutes
	app.Post("/tests/createClients", routes.CreateClient)
	app.Get("/tests/listClients", routes.ListClients)
	app.Get("/tests/getClient/:id", routes.GetClient)
	app.Put("tests/updateClient/:id", routes.UpdateClient)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":8000")
}
