package book

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/calibre"
)

const maxBookQuery = 30

func index(ctx *fiber.Ctx) error {
	random, err := getQueryInt(ctx, "random", 8, maxBookQuery)
	if err != nil {
		return ctx.JSON(map[string]string{"err": "exception", "msg": err.Error()})
	}

	recent, err := getQueryInt(ctx, "recent", 10, maxBookQuery)
	if err != nil {
		return ctx.JSON(map[string]string{"err": "exception", "msg": err.Error()})
	}

	// Query a random book id set.
	randomIDs, err := calibre.QueryRandomBookIDs(ctx.UserContext(), random)
	if err != nil {
		return ctx.JSON(map[string]string{"err": "exception", "msg": err.Error()})
	}

	// Query the random books by ids.
	randoms, err := calibre.QueryBooksByIDs(ctx.UserContext(), randomIDs)
	if err != nil {
		return ctx.JSON(map[string]string{"err": "exception", "msg": err.Error()})
	}

	// Query recent books.
	// SELECT * FROM table ORDER BY id DESC LIMIT x)
	news, err := calibre.QueryBooks(ctx.UserContext(), 1, recent)
	if err != nil {
		return ctx.JSON(map[string]string{"err": "exception", "msg": err.Error()})
	}

	return ctx.JSON(map[string]any{
		"random_books_count": len(randoms),
		"new_books_count":    len(news),
		"random_books":       randoms,
		"new_books":          news,
		"msg":                "ok",
	})
}

func getQueryInt(ctx *fiber.Ctx, name string, dv, max int) (int, error) {
	query := ctx.Query(name, "")
	if query == "" {
		return dv, nil
	}
	val, err := strconv.Atoi(query)
	if err != nil {
		return 0, err
	}

	if val < max {
		return val, nil
	} else {
		return max, nil
	}
}
