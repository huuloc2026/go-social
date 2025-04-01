package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/interfaces/http/handlers"
	"github.com/huuloc2026/go-social/internal/interfaces/http/middlewares"
)

func SetupRoutes(app *fiber.App, authUseCase usecases.AuthUseCase, userUseCase usecases.UserUseCase) {
	authController := handlers.NewAuthController(authUseCase)
	userController := handlers.NewUserController(userUseCase)

	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	// User routes
	user := app.Group("/users")
	user.Use(middlewares.AuthMiddleware(authUseCase))
	user.Get("/:id", userController.GetUser)
	user.Put("/:id", userController.UpdateUser)
	user.Delete("/:id", userController.DeleteUser)
}
