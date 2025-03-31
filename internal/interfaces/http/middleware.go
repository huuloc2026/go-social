package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/pkg/utils"
)

// Middleware kiểm tra JWT Token
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		c.Locals("userID", claims.UserID)
		return c.Next()
	}
}
