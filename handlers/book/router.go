package book

import "github.com/gofiber/fiber/v2"

// Handlers will add the handlers to fiber.
func Handlers(app *fiber.App) {
	app.Get("/api/index", index)
}
