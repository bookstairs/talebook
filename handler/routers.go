package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/handler/admin"
	"github.com/bookstairs/talebook/handler/book"
	"github.com/bookstairs/talebook/handler/file"
	"github.com/bookstairs/talebook/handler/metadata"
	"github.com/bookstairs/talebook/handler/opds"
	"github.com/bookstairs/talebook/handler/scan"
	"github.com/bookstairs/talebook/handler/user"
)

// registerHandlers will add the handlers to fiber.
func registerHandlers(app *fiber.App) {
	admin.Handlers(app)
	book.Handlers(app)
	file.Handlers(app)
	metadata.Handlers(app)
	opds.Handlers(app)
	scan.Handlers(app)
	user.Handlers(app)
}
