package main

import (

	// Import the generated templ file (replace with your actual module path)

	handlers "github.com/AramisAra/BravusServer/handlers"
	viewrenders "github.com/AramisAra/BravusServer/handlers/viewhandlers"
	middlewares "github.com/AramisAra/BravusServer/middleware"
	"github.com/AramisAra/BravusServer/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func testRoutes(app *fiber.App) {
	app.Get("/protected", middlewares.Protected(), handlers.Protected)
}

func viewroutes(app *fiber.App) {
	app.Group("/userviews")

	app.Group("/ownerviews")

	app.Group("/mainviews")
	app.Get("/", viewrenders.LandingPage)
}

func databaseRoutes(app *fiber.App) {
	// Client Routes
	client := app.Group("/client")
	client.Post("/create", handlers.RegisterClient)
	client.Post("/login", handlers.LoginClient)
	client.Get("/Get", handlers.ListClient)
	client.Get("/get", handlers.GetClient)
	client.Put("/update", handlers.UpdateClient)
	client.Delete("/delete", handlers.DeleteClient)

	// Owner Routes
	owner := app.Group("/owner")
	owner.Post("/create", handlers.RegisterOwner)
	owner.Post("/login", handlers.LoginOwner)
	owner.Get("/Get", handlers.ListOwner)
	owner.Get("/get", handlers.GetOwner)
	owner.Put("/update", handlers.UpdateOwner)
	owner.Delete("/delete", handlers.DeleteOwner)

	// Service Routes
	service := app.Group("/service")
	service.Post("/create", handlers.CreateService)
	service.Get("/Get", handlers.ListService)
	service.Put("/update", handlers.UpdateService)
	service.Delete("/delete", handlers.DeleteService)
}

func main() {

	// This engine is the instance of the template
	// controller
	engine := jet.New("./view", ".jet")

	// This gets the .env files
	utils.GetDot()

	// Config of the fiber app
	server := fiber.New(fiber.Config{
		Views:         engine,
		CaseSensitive: true,
	})

	// These are the routes to database
	databaseRoutes(server)
	// These are the jwt test routes
	testRoutes(server)
	// This router manages the view of the applications
	viewroutes(server)

	server.Listen(":8010")
}
