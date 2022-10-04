package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/config"
	"github.com/bookstairs/talebook/handler/book"
	"github.com/bookstairs/talebook/handler/file"
	"github.com/bookstairs/talebook/handler/user"
)

// manually add the handlers here.
func registerHandlers(app *fiber.App) {
	// Admin Handlers Checklist
	//
	// /api/admin/ssl
	// /api/admin/users
	// /api/admin/install
	// /api/admin/settings
	// /api/admin/testmail
	// /api/admin/book/list

	// Book Handlers Checklist
	//
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
	app.Get("/api/index", book.Index)
	app.Get("/api/book/:id<int>", book.GetBookByID)
	app.Get("/api/recent", book.ListAllBook)
	app.Get("/api/book/:id<int>.:ext", book.DownloadBookByID)

	// File Handlers
	app.Get("/get/pcover", file.ProxyCover)
	app.Get("/get/progress/:bid<int>", file.Progress)
	app.Get("/get/extract/:bid<int>", file.Extract)
	app.Get("/get/:kind/:id<int>.jpg", file.ImageCover)

	// Metadata Handlers Checklist
	//
	// /api/(author|publisher|tag|rating|series)
	// /api/(author|publisher|tag|rating|series)/(.*)
	// /api/author/(.*)/update
	// /api/publisher/(.*)/update

	// OPDS Handlers Checklist
	//
	// /opds/?
	// /opds/nav/(.*)
	// /opds/category/(.*)/(.*)
	// /opds/categorygroup/(.*)/(.*)
	// /opds/search/(.*)

	// Scan Handlers Checklist
	//
	// /api/admin/scan/list
	// /api/admin/scan/run
	// /api/admin/scan/status
	// /api/admin/scan/delete
	// /api/admin/scan/mark
	// /api/admin/import/run
	// /api/admin/import /status

	// Use Handlers Checklist
	//
	// /api/welcome
	// /api/user/sign_in
	// /api/user/sign_up
	// /api/user/sign_out
	// /api/user/update
	// /api/user/reset
	// /api/user/active/send
	// /api/active/(.*)/(.*)
	// /api/done/
	app.Get("/api/user/info", user.Info)
	app.Get("/api/user/messages", user.GetMessages)
}

// If you have some dynamically init logic, add it here.
func initHandlers(c *config.ServerConfig) {
	if err := file.InitCoverCache(c); err != nil {
		log.Fatalln(err)
	}
}
