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
	app.Post("/dev/createClientwithAnimal", routes.CreateClientAndAnimal)
	// Client's dataRoutes
	app.Post("/dev/createAnimal/:uuid", routes.CreateAnimal)
	app.Get("/dev/listClients", routes.ListClients)
	app.Get("/dev/getClient/:uuid", routes.GetClient)
	app.Put("/dev/updateClient/:uuid", routes.UpdateClient)
	app.Delete("/dev/deleteClient/:uuid", routes.DeleteClient)
	app.Get("/dev/getAppointment/:uuid", routes.GetAppointment)
	// Service's dataRoutes
	app.Post("/dev/createServices", routes.CreateService)
	app.Get("/dev/listServices", routes.ListService)
	app.Get("/dev/getService/:uuid", routes.GetService)
	app.Put("/dev/updateSevice/:uuid", routes.UpdateService)
	app.Delete("/dev/deleteService/:uuid", routes.DeleteService)
	// Owner's dataRoutes
	app.Post("/dev/createOwner", routes.CreateOwner)
	app.Get("/dev/listOwners", routes.ListOwners)
	app.Get("/dev/getOwner/:uuid", routes.GetOwner)
	app.Put("/dev/updateOwner/:uuid", routes.UpdateOwner)
	app.Delete("/dev/deleteOwner/:uuid", routes.DeleteOwner)
	app.Get("/dev/getAppointmentOwner/:uuid", routes.GetAppointmentOwner)
	//Animal's dataRoutes
	app.Get("/dev/listAnimals", routes.ListAnimals)
	app.Get("/dev/getAnimal/:uuid", routes.GetAnimal)
	app.Put("/dev/updateAnimal/:uuid", routes.UpdateAnimal)
	// Appointment's dataroutes --- This include the animal appointments routes too
	app.Post("/dev/createAppointment/:uuidUser/:uuidAnimal/:uuidOwner", routes.CreateAppointment)
	app.Put("/dev/updateAppointment/:uuid", routes.UpdateAppointment)
	app.Delete("/dev/deleteAppointment/:uuid", routes.DeleteAppointment)
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
