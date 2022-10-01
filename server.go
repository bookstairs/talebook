package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/handlers"
)

type ServerConfig struct {
	Port        int    `yaml:"port"`        // The binding port for backend server.
	WorkingPath string `yaml:"workingPath"` // The working directory for talebook.
	LibraryPath string `yaml:"libraryPath"` // The calibre library directory.
	EncryptKey  string `yaml:"encryptKey"`  // This is used to encrypt the cookie.
	Limit       int    `yaml:"limit"`       // Allowed request per seconds.
	CalibreDB   string `yaml:"calibreDB"`   // The executable file calibredb for adding books.
	Convert     string `yaml:"convert"`     // The executable file ebook-convert for converting books.
	Debug       bool   `yaml:"debug"`       // Enable debug log and metrics monitor and anything else.
}

func (c *ServerConfig) GetPath(paths ...string) string {
	return filepath.Join(c.WorkingPath, filepath.Join(paths...))
}

func DefaultSeverConfig() *ServerConfig {
	// Init the config variables with some default values.
	w, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return &ServerConfig{
		Port:        8000,
		WorkingPath: w,
		LibraryPath: filepath.Join(w, "library"),
		EncryptKey:  "this-is-an-encrypt-key",
		Limit:       100,
		CalibreDB:   calibre.DefaultCalibreDB,
		Convert:     calibre.DefaultConvert,
		Debug:       false,
	}
}

// StartServer will start the talebook server.
func StartServer(c *ServerConfig) {
	// Create working directories.
	createWorkingPaths(c)

	// Create fiber application.
	app := fiber.New(fiber.Config{
		AppName:                 "Talebook",
		Immutable:               true,
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
	app.Use(encryptcookie.New(encryptcookie.Config{Key: c.EncryptKey}))

	// Add ETag.
	app.Use(etag.New())

	// Add ratelimit for avoiding spiders and anything else like the book downloader.
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        c.Limit * 30,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			// We will drop the request when it exceeds the limits.
			return c.SendStatus(403)
		},
	}))

	if c.Debug {
		// Add a metrics monitor.
		app.Get("/metrics", monitor.New(monitor.Config{Title: "Talebook Monitor"}))
	}

	// Add API backend.
	handlers.RegisterHandlers(app)

	// Allow end user to add custom frontend files.
	// This would override the default frontend files. Use it at your own risk.
	app.Static("/", c.GetPath("statics"))

	// The frontend application.
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(frontend),
		PathPrefix:   "app/dist",
		Browse:       false,
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
	}))

	// Listen on given port.
	log.Fatal(app.Listen(":" + strconv.Itoa(c.Port)))
}

// createWorkingPaths will create all the required working directory.
func createWorkingPaths(c *ServerConfig) {
	// Internal method for making all the directories if it's not existed.
	createPath := func(path string) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	// Internal method for checking if a file exists.
	fileExist := func(path string) bool {
		_, err := os.Stat(path)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			// We can't access this file because it's unreachable.
			log.Fatal(err)
		}
		return err == nil
	}

	// Create all the working directories if they are not existed.
	createPath(c.GetPath("statics"))
	createPath(c.LibraryPath)

	// Extract the default calibre library in case of failure.
	if !fileExist(calibre.GetDatabase(c.LibraryPath)) {
		log.Print("No calibre library found extract the default library.")
		err := extractDefaultLibrary(c.LibraryPath)
		if err != nil {
			// Use log.Fatal may not execute the defer method.
			log.Fatal(err)
		}
	}
}
