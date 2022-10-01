package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/handlers/admin"
	"github.com/bookstairs/talebook/handlers/book"
	"github.com/bookstairs/talebook/handlers/file"
	"github.com/bookstairs/talebook/handlers/metadata"
	"github.com/bookstairs/talebook/handlers/opds"
	"github.com/bookstairs/talebook/handlers/scan"
	"github.com/bookstairs/talebook/handlers/user"
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
