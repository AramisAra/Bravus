package viewrenders

import "github.com/gofiber/fiber/v2"

func LandingPage(c *fiber.Ctx) error {
	// This renders the landing pages
	return c.Render("LandingPage", fiber.Map{
		"Title": "Bravus",
	})
}
