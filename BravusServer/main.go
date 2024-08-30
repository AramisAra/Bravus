package main

import (
	handlers "github.com/AramisAra/BravusServer/handlers"
	middlewares "github.com/AramisAra/BravusServer/middleware"
	"github.com/AramisAra/BravusServer/utils"
	"github.com/gofiber/fiber/v2"
)

func testRoutes(app *fiber.App) {
	app.Get("/protected", middlewares.Protected(), handlers.Protected)
}

func databaseRoutes(app *fiber.App) {
	// Client Routes
	client := app.Group("/client")
	client.Post("/create", handlers.RegisterClient)
	client.Post("/login", handlers.LoginClient)
	// List all Clients
	client.Get("/Get", handlers.ListClient)
	// Gets Client based on ID
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
}

func main() {
	utils.GetDot()
	server := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	// Backend home get
	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Backend")
	})

	databaseRoutes(server)
	testRoutes(server)

	server.Listen(":8010")
}
