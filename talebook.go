package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	app := fiber.New()

	// The frontend application.
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(bundle),
		PathPrefix:   "app/dist",
		Browse:       false,
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
	}))

	app.Listen(":8000")
}
