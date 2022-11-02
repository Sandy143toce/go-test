package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// func setupRoutes(app *fiber.App) {
// 	app.Get("/", handlers.Home)
// }

func Home(c *fiber.Ctx) error {
	return c.SendString("Webpage hai yee with air")
}