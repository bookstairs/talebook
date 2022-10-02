package file

import "github.com/gofiber/fiber/v2"

// Handlers will add the handlers to fiber.
//
// /get/pcover
// /get/progress/([0-9]+)
// /get/extract/(.*)
// /get/(.*)/(.*)
func Handlers(app *fiber.App) {
	// TODO Add handlers for this endpoint.
	app.Get("/get/pcover", proxyCover)
	app.Get("/get/progress/:bid<int>", progress)
	app.Get("/get/extract/:bid<int>", extract)
	app.Get("/get/:fmt/:bid<int>", load)
}
