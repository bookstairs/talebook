package book

import "github.com/gofiber/fiber/v2"

// Handlers will add the handlers to fiber.
//
// /api/index
// /api/search
// /api/recent
// /api/hot
// /api/book/nav
// /api/book/upload
// /api/book/([0-9]+)
// /api/book/([0-9]+)/delete
// /api/book/([0-9]+)/edit
// /api/book/([0-9]+)\.(.+)
// /api/book/([0-9]+)/push
// /api/book/([0-9]+)/refer
// /read/([0-9]+)
func Handlers(app *fiber.App) {
	app.Get("/api/index", index)
}
