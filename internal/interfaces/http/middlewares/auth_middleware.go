package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/utils"
)

func AuthMiddleware(authUseCase usecases.AuthUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponse(ctx, fiber.StatusUnauthorized, "Missing authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return utils.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid token format")
		}

		userID, err := authUseCase.ValidateToken(tokenString)
		if err != nil {
			return utils.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid or expired token")
		}

		ctx.Locals("userID", userID)
		return ctx.Next()
	}
}
