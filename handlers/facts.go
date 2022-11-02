package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Sandy143toce/go-test/models"
	"github.com/Sandy143toce/go-test/database"
)

// func setupRoutes(app *fiber.App) {
// 	app.Get("/", handlers.Home)
// }

func Home(c *fiber.Ctx) error {
	return c.SendString("Webpage hai yee with air")
}