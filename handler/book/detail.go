package book

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/handler/common"
)

func GetBookByID(ctx *fiber.Ctx) error {
	id, err := common.GetParamInt(ctx, "id")
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	book, err := calibre.QueryBookDetailByID(ctx.UserContext(), id)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	return common.SuccResp(ctx, map[string]any{
		// TODO Support Kindle sender
		"kindle_sender": "sender@talebook.org",
		"book":          book,
	})
}
