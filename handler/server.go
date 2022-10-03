package handler

import (
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/bookstairs/talebook/config"
)

// StartServer will start the talebook server.
func StartServer(c *config.ServerConfig) {
	// Create fiber application.
	app := fiber.New(fiber.Config{
		AppName:                 "Talebook",
		Immutable:               true,
		EnableTrustedProxyCheck: true,
		JSONEncoder:             json.Marshal,
		JSONDecoder:             json.Unmarshal,
	})

	// Add cache support. We will disable cache in debug mode for development purpose.
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true" || c.IP() == "127.0.0.1"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	// Encrypt the cookie for end user.
	app.Use(encryptcookie.New(encryptcookie.Config{Key: c.EncryptKey}))

	// Add ETag.
	app.Use(etag.New())

	// Add ratelimit for avoiding spiders and anything else like the book downloader.
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			// We only cache the API request.
			return c.IP() == "127.0.0.1" || !strings.HasPrefix(c.OriginalURL(), "/api")
		},
		Max:        c.Limit * 30,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			// We will drop the request when it exceeds the limits.
			return c.SendStatus(403)
		},
	}))

	// Add a metrics monitor.
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Talebook Monitor"}))

	// Add API backend.
	registerHandlers(app)

	// Initialize all the API handlers.
	initHandlers(c)

	// Allow end user to add custom frontend files.
	// This would override the default frontend files. Use it at your own risk.
	app.Static("/", c.GetPath("statics"))

	// The frontend application.
	statics, err := fs.Sub(c.Frontend, "app/dist")
	if err != nil {
		log.Fatal(err)
	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(statics),
		Browse:       false,
		Index:        "index.html",
		NotFoundFile: "index.html", // We should keep this file as the not found file for frontend SPA.
		MaxAge:       3600,
	}))

	// Listen on given port.
	log.Fatal(app.Listen(":" + strconv.Itoa(c.Port)))
}
