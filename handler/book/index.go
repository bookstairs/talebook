package book

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/handler/common"
)

const maxBookQuery = 30

func Index(ctx *fiber.Ctx) error {
	random, err := common.GetQueryInt(ctx, "random", 8, maxBookQuery)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	recent, err := common.GetQueryInt(ctx, "recent", 10, maxBookQuery)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	// Query a random book id set.
	randomIDs, err := calibre.QueryRandomBookIDs(ctx.UserContext(), random)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	// Query the random books by ids.
	randoms, err := calibre.QueryBooksByIDs(ctx.UserContext(), randomIDs)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	// Query recent books.
	// SELECT * FROM table ORDER BY id DESC LIMIT x)
	news, err := calibre.QueryBooks(ctx.UserContext(), 1, recent)
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	return common.SuccResp(ctx, map[string]any{
		"random_books_count": len(randoms),
		"new_books_count":    len(news),
		"random_books":       randoms,
		"new_books":          news,
	})
}
