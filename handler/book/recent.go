package book

import (
	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/handler/common"
	"github.com/gofiber/fiber/v2"
)

func ListAllBook(ctx *fiber.Ctx) error {
	start, err := common.GetQueryInt(ctx, "start", 0, 99999999999)
	if err != nil {
		return common.ErrResp(ctx, err)
	}
	size, err := common.GetQueryInt(ctx, "size", 60, 100)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	books, total, err := calibre.ListBookByPage(ctx.UserContext(), start, size)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	return common.SuccResp(ctx, map[string]any{
		// TODO Support Kindle sender
		"title": "新书推荐",
		"books": books,
		"total": total,
	})
}
