package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	// Create fiber application.
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
	})

	// Add cache support.
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	// Encrypt the cookie for end user.
	app.Use(encryptcookie.New(encryptcookie.Config{
		// FIXME Change it to a configurable constant.
		Key: "this-is-an-encrypt-key",
	}))

	// Add ETag.
	app.Use(etag.New())

	// Add ratelimit for avoiding spiders and anything else like the book downloader.
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(403)
		},
	}))

	// Add a metrics monitor. FIXME This should be enabled with debug purpose.
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Talebook Monitor"}))

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
