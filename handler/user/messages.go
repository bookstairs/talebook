package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/handler/common"
	"github.com/bookstairs/talebook/model"
)

// GetMessages would serve /api/user/messages
func GetMessages(ctx *fiber.Ctx) error {
	// TODO Implement this method in the future.
	return common.SuccResp(ctx, map[string]any{
		"message": []model.Message{},
	})
}
