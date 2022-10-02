package file

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/handler/common"
)

func Progress(ctx *fiber.Ctx) error {
	id, _ := common.GetParamInt(ctx, "bid") // This must be an int value.
	formats, err := calibre.QueryBookFormats(ctx.UserContext(), id)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	for _, b := range formats {
		if b.Format == "EPUB" {
			if _, err := os.Stat(b.Path); err != nil {
				return common.NotFound(ctx, err)
			}
			return ctx.SendFile(b.Path)
		}
	}

	// TODO Support the PDF file by the same time.
	return common.ErrResp(ctx, fmt.Errorf("couldn't found the epub file"))
}
