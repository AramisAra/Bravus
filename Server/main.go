package main

import (
	"log"
	"os"

	database "github.com/AramisAra/GroomingApp/database"
	handlers "github.com/AramisAra/GroomingApp/handlers"
	"github.com/AramisAra/GroomingApp/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func dbhandlers(app *fiber.App) {
	// HealthCheck
	app.Get("/health", HealthCheck)
	app.Post("/login", handlers.LoginClient)
	app.Post("/loginowner", handlers.LoginOwner)

	// Tests Functions
	app.Post("/dev/registerClient", handlers.RegisterClient)
	// Client's datahandlers
	app.Post("/dev/createAnimal/:uuid", handlers.CreateAnimal)
	app.Get("/dev/listClients", handlers.ListClients)
	app.Get("/dev/getClient/:uuid", handlers.GetClient)
	app.Put("/dev/updateClient/:uuid", handlers.UpdateClient)
	app.Delete("/dev/deleteClient/:uuid", handlers.DeleteClient)
	app.Get("/dev/getAppointment/:uuid", handlers.GetAppointment)
	// Service's datahandlers
	app.Post("/dev/createServices", handlers.CreateService)
	app.Get("/dev/listServices", handlers.ListService)
	app.Get("/dev/getService/:uuid", handlers.GetService)
	app.Put("/dev/updateSevice/:uuid", handlers.UpdateService)
	app.Delete("/dev/deleteService/:uuid", handlers.DeleteService)
	// Owner's datahandlers
	app.Post("/dev/registerOwner", handlers.RegisterOwner)
	app.Get("/dev/listOwners", handlers.ListOwners)
	app.Get("/dev/getOwner/:uuid", handlers.GetOwner)
	app.Put("/dev/updateOwner/:uuid", handlers.UpdateOwner)
	app.Delete("/dev/deleteOwner/:uuid", handlers.DeleteOwner)
	app.Get("/dev/getAppointmentOwner/:uuid", handlers.GetAppointmentOwner)
	//Animal's datahandlers
	app.Get("/dev/listAnimals", handlers.ListAnimals)
	app.Get("/dev/getAnimal/:uuid", handlers.GetAnimal)
	app.Put("/dev/updateAnimal/:uuid", handlers.UpdateAnimal)
	// Appointment's datahandlers --- This include the animal appointments handlers too
	app.Post("/dev/createAppointment/:uuidUser/:uuidAnimal/:uuidOwner", handlers.CreateAppointment)
	app.Put("/dev/updateAppointment/:uuid", handlers.UpdateAppointment)
	app.Delete("/dev/deleteAppointment/:uuid", handlers.DeleteAppointment)
}

func loghandlers(jwt fiber.Handler, app *fiber.App) {
	app.Get("/protected", jwt, handlers.Protected)
	app.Get("/dev/sheets", jwt, handlers.GetSheets)
	app.Get("/dev/table/:sheet", jwt, handlers.GetTable)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the Env: ", err)
	}
	// This is how the database connects
	dsn := os.Getenv("DSN")
	database.ConnectDb(dsn)

	// This is the main overall the app_api
	app := fiber.New()
	jwt := middlewares.NewAuthMiddleware()
	loghandlers(jwt, app)
	dbhandlers(app)
	app.Listen(":8000")
}
