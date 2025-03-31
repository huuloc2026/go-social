package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/huuloc2026/go-social/internal/config"
	"github.com/huuloc2026/go-social/internal/interfaces/adapters/http/handlers"
	"github.com/huuloc2026/go-social/internal/usecases"
)

func SetupRoutes(app *fiber.App, userUsecase usecases.UserUsecase, cfg *config.Config) {
	// Middleware
	app.Use(logger.New())
	app.Use(NewAuthMiddleware(cfg.App.JWTSecret).Middleware)

	// Handlers
	userHandler := handlers.NewUserHandler(userUsecase)

	// Routes
	api := app.Group("/api")

	// Auth routes (no auth required)
	auth := api.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	// Authenticated routes
	api.Get("/me", userHandler.GetCurrentUser)
}
