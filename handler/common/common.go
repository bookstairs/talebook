package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var handlerRegistry map[string]map[string][]fiber.Handler

// RegisterHandlers will add all the handles into registry.
func RegisterHandlers(method, path string, handler fiber.Handler) {
	m := handlerRegistry[method]
	if m == nil {
		m = make(map[string][]fiber.Handler, 0)
	}
	hs := m[path]
	hs = append(hs, handler)
	m[path] = hs
	handlerRegistry[method] = m
}

// SuccResp will return the talebook supported success response.
func SuccResp(ctx *fiber.Ctx, data map[string]any) error {
	data["msg"] = ""
	data["err"] = "ok"
	return ctx.JSON(data)
}

// ErrResp will return the talebook supported error response.
func ErrResp(ctx *fiber.Ctx, err error) error {
	return ctx.JSON(map[string]any{"err": "exception", "msg": err})
}

func GetQueryInt(ctx *fiber.Ctx, name string, dv, max int) (int, error) {
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
