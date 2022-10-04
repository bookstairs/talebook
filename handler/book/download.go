package book

import (
	"fmt"
	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/handler/common"
	"github.com/bookstairs/talebook/model"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func DownloadBookByID(ctx *fiber.Ctx) error {
	id, err := common.GetParamInt(ctx, "id")
	if err != nil {
		return common.ErrResp(ctx, err)
	}
	ext, err := common.GetParamString(ctx, "ext")
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	book, err := calibre.QueryBookDetailByID(ctx.UserContext(), id)
	if err != nil {
		return common.ErrResp(ctx, err)
	}
	for _, file := range book.Files {
		if strings.EqualFold(ext, file.Format) {
			return ctx.SendFile(getFilePath(book))
		}
	}
	return common.ErrResp(ctx, fmt.Errorf("not found book %d.%s", id, ext))
}

func getFilePath(book *model.Book) string {
	//TODO need get a file path
	return ""
}
