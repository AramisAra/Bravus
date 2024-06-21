package main

import (
	"log"
	"os"

	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/handlers"
	middlewares "github.com/AramisAra/BravusBackend/middleware"
	"github.com/AramisAra/BravusBackend/sheetsapi"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
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
	// Service Routes
	service := app.Group("/service")
	service.Post("/create/:uuid", handlers.CreateService)
	service.Get("/get", handlers.ListService)
	service.Put("/update/:uuid", handlers.UpdateService)
	service.Delete("/delete/:uuid", handlers.DeleteService)
	// Owner Routes
	owner := app.Group("/owner")
	owner.Get("/get", handlers.ListOwners)
	owner.Get("/get/:uuid", handlers.GetOwner)
	owner.Put("/update/:uuid", handlers.UpdateOwner)
	owner.Delete("/delete/:uuid", handlers.DeleteOwner)
}

func SheetsHandler(app *fiber.App) {
	//jwt := middlewares.NewAuthMiddleware()
	sheetapi := app.Group("/sheetapi")
	sheetapi.Get("/auth", sheetsapi.AuthGoogle)
	sheetapi.Get("/auth/callback", sheetsapi.AuthCallback)
	sheetapi.Post("/createSheet", sheetsapi.CreateSheet)
	sheetapi.Get("/sheet", sheetsapi.GetSheet)
	// Get Values will return  a default of 1500 Cells but it only return the filled cells
	sheetapi.Get("/getValues", sheetsapi.GetSheetValues)
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
	sheetsapi.Start()
	engine := django.New("./views", ".django")

	// This is the main overall the app_api
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	jwt := middlewares.NewAuthMiddleware()
	LoginSystem(app)
	DatabaseHandlers(app)
	SheetsHandler(app)
	app.Get("/protected", jwt, handlers.Protected)
	app.Listen(":8000")
}
