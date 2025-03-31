package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/interfaces/http/handlers"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	api := app.Group("/api")

	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	// Protected routes
	user := api.Group("/user", JWTMiddleware())
	user.Get("/:id", userHandler.GetUserByID)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
