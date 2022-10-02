package file

import (
	"github.com/bookstairs/talebook/calibre"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

func progress(ctx *fiber.Ctx) error {
	bookID, _ := strconv.ParseInt(ctx.Params("bid"), 10, 0)

	bookFormats, err := calibre.BookFormatsQuery(ctx.UserContext(), bookID)
	if err != nil {
		return ctx.JSON(map[string]string{"err": "exception", "msg": err.Error()})
	}
	for _, b := range bookFormats {
		if b.Format == "EPUB" {

			_, err := os.Stat(b.Path)
			if err != nil {
				return ctx.Status(404).JSON(map[string]string{"err": "exception", "msg": err.Error()})
			}
			// TODO unzip epub
			return ctx.SendFile(b.Path)
		}
	}
	return ctx.JSON(map[string]string{"err": "exception", "msg": "not found epub file"})
}
