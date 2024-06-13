package main

import (
	"os"

	database "github.com/AramisAra/GroomingApp/database"
	routes "github.com/AramisAra/GroomingApp/database/routes"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func setupRoutes(app *fiber.App) {
	// HealthCheck
	app.Get("/health", HealthCheck)
	// Tests Functions
	app.Post("/tests/createClientwithAnimal", routes.CreateClientAndAnimal)
	// Client's dataRoutes
	app.Post("/tests/createAnimal/:uuid", routes.CreateAnimal)
	app.Get("/tests/listClients", routes.ListClients)
	app.Get("/tests/getClient/:uuid", routes.GetClient)
	app.Put("tests/updateClient/:uuid", routes.UpdateClient)
	app.Delete("tests/deleteClient/:uuid", routes.DeleteClient)
	// Service's dataRoutes
	app.Post("/tests/createServices", routes.CreateService)
	app.Get("/tests/listServices", routes.ListService)
	app.Get("/tests/getService/:uuid", routes.GetService)
	app.Put("/tests/updateSevice/:uuid", routes.UpdateService)
	app.Delete("/tests/deleteService/:uuid", routes.DeleteService)
	// Owner's dataRoutes
	app.Post("/tests/createOwner", routes.CreateOwner)
	app.Get("/tests/listOwners", routes.ListOwners)
	app.Get("/tests/getOwner/:uuid", routes.GetOwner)
	app.Put("/tests/updateOwner/:uuid", routes.UpdateOwner)
	app.Delete("/tests/deleteOwner/:uuid", routes.DeleteOwner)
	//Animal's dataRoutes
	app.Get("/tests/listAnimals", routes.ListAnimals)
	app.Get("/tests/getAnimal/:uuid", routes.GetAnimal)
	app.Put("/tests/updateAnimal/:uuid", routes.UpdateAnimal)
	// Appointment's dataroutes --- This include the animal appointments routes too

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
