package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Sandy143toce/go-test"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}