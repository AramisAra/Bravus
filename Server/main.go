package main

import (
	"os"

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
	app.Delete("tests/deleteClient/:id", routes.DeleteClient)
	// Service's dataRoutes
	app.Post("/tests/createServices", routes.CreateService)
	app.Get("/tests/listServices", routes.ListService)
	app.Get("/tests/getService/:id", routes.GetService)
	app.Put("/tests/updateSevice/:id", routes.UpdateService)
	app.Delete("/tests/deleteService/:id", routes.DeleteClient)
	// Appointment's dataRoutes

}

func main() {
	// This is how the database connects
	dsn := os.Getenv("dsn")
	database.ConnectDb(dsn)

	// This is the main overall the app_api
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":8000")
}
