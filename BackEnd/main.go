package main

import (
	"log"
	"os"

	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/googleapis"
	"github.com/AramisAra/BravusBackend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func LoginSystem(app *fiber.App) {
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
	client.Get("/get", handlers.GetClient)
	client.Get("/get", handlers.ListClients)
	client.Put("/update", handlers.UpdateClient)
	client.Delete("/delete", handlers.DeleteClient)
	// Animal Routes
	animal := app.Group("/animal")
	animal.Post("/create", handlers.CreateAnimal)
	animal.Delete("/delete", handlers.DeleteAnimal)
	animal.Put("/update", handlers.UpdateAnimal)
	// Appointment Routes
	appointment := app.Group("/appointment")
	appointment.Post("/create", handlers.CreateAppointment)
	appointment.Get("/getforclient", handlers.GetAppointmentClient)
	appointment.Get("/getforowner", handlers.GetAppointmentOwner)
	appointment.Delete("/delete", handlers.DeleteAppointment)
	appointment.Put("/update", handlers.UpdateAppointment)
	// Service Routes
	service := app.Group("/service")
	service.Post("/create", handlers.CreateService)
	service.Get("/get", handlers.ListService)
	service.Put("/update", handlers.UpdateService)
	service.Delete("/delete", handlers.DeleteService)
	// Owner Routes
	owner := app.Group("/owner")
	owner.Get("/get", handlers.ListOwners)
	owner.Get("/get", handlers.GetOwner)
	owner.Put("/update", handlers.UpdateOwner)
	owner.Delete("/delete", handlers.DeleteOwner)
}

func SheetsHandler(app *fiber.App) {
	sheetapi := app.Group("/sheetapi")
	sheetapi.Get("/auth", googleapis.AuthGoogle)
	sheetapi.Get("/auth/callback", googleapis.AuthCallback)
	sheetapi.Post("/createSheet", googleapis.CreateSheet)
	sheetapi.Get("/getSheet", googleapis.GetSheet)
	// Get Values will return  a default of 1500 Cells but it only return the filled cells
	sheetapi.Get("/getValues", googleapis.GetSheetValues)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the Env: ", err)
	}
	// This is how the database connects
	dsn := os.Getenv("DSN")
	database.ConnectDb(dsn)

	// Setting auth to google servers
	googleapis.Start()

	// This is the main overall the app_api
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://172.24.195.132:3000, https://aramisara.github.io, http://34.204.43.154:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	LoginSystem(app)
	DatabaseHandlers(app)
	SheetsHandler(app)

	// Start server
	log.Fatal(app.Listen(":8000"))
}
