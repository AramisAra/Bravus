package main

import (
	handlers "github.com/AramisAra/BravusServer/handlers"
	"github.com/AramisAra/BravusServer/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.GetDot()
	server := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Backend")
	})

	server.Get("/Get", handlers.ListClient)
	server.Get("/get", handlers.GetClient)
	server.Put("/update", handlers.UpdateClient)
	server.Delete("/delete", handlers.DeleteClient)
	server.Post("/create", handlers.RegisterClient)
	server.Post("/login", handlers.LoginClient)
	server.Listen(":8010")
}
