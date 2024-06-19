package main

import (
	"log"
	"os"

	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func loginSystem(app *fiber.App) {
	// login system
	login := app.Group("/login")
	login.Post("/Rowner", handlers.RegisterOwner)
	login.Post("/Rclient", handlers.RegisterClient)
	login.Post("/Lowner", handlers.LoginOwner)
	login.Post("/Lclient", handlers.LoginClient)
}
func DatabaseHandlers(app *fiber.App) {
	// Client Routes
	client := app.Group("/client")
	client.Get("/get/:uuid", handlers.GetClient)
	client.Get("/get", handlers.ListClients)
	client.Put("/update/:uuid", handlers.UpdateClient)
	client.Delete("/delete/:uuid", handlers.DeleteClient)
	// Animal Routes
	animal := app.Group("/animal")
	animal.Post("/create/:uuid", handlers.CreateAnimal)
	animal.Delete("/delete/:uuid", handlers.DeleteAnimal)
	animal.Put("/update/:uuid", handlers.UpdateAnimal)
	// Appointment Routes
	appointment := app.Group("/appointment")
	appointment.Post("/create/:uuid/:uuid/uuid", handlers.CreateAppointment)
	appointment.Get("/get/:uuid", handlers.GetAppointment)
	appointment.Delete("/delete/:uuid", handlers.DeleteAppointment)
	appointment.Put("/update/:uuid", handlers.UpdateAppointment)
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
	loginSystem(app)
	DatabaseHandlers(app)

	app.Listen(":8000")
}
