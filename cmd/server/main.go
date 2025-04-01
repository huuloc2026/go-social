package main

import (
	"log"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/config"
	"github.com/huuloc2026/go-social/infrastructure/database"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/interfaces/http"
	"github.com/huuloc2026/go-social/internal/interfaces/http/repositories"
	"github.com/huuloc2026/go-social/internal/utils"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Initialize JWT
	utils.InitJWT(cfg.JWTSecret)

	// Initialize database
	db, err := database.NewPostgresDB(&cfg)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo)
	userUseCase := usecases.NewUserUseCase(userRepo)

	// Create Fiber app
	app := fiber.New()

	// Setup routes
	http.SetupRoutes(app, authUseCase, userUseCase)

	// Start server
	port := ":" + cfg.Port
	if err := app.Listen(port); err != nil {
		log.Fatal("Failed to start server:", err)
		os.Exit(1)
	}
}
