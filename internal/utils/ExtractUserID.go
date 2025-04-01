package utils

import (
	"github.com/gofiber/fiber/v2"
	domainErrors "github.com/huuloc2026/go-social/internal/domain/errors"
)

func ExtractUserID(ctx *fiber.Ctx) (uint, error) {
	userID, ok := ctx.Locals("userID").(uint)
	if !ok {
		return 0, domainErrors.ErrUnauthorized
	}
	return userID, nil
}
