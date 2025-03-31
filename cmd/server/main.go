package main

import (
	"log"

	"github.com/huuloc2026/go-social/internal/interfaces/http"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/config"
	"github.com/huuloc2026/go-social/internal/config/database"
	"github.com/huuloc2026/go-social/internal/domain"
	"github.com/huuloc2026/go-social/internal/usecases"
)

func main() {
	// Load configuration
	cfg := config.Load()

	//Connect DB
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Auto migrate models
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repositories
	userRepo := domain.NewUserRepository(db.DB)

	// Initialize use cases
	userUsecase := usecases.NewUserUsecase(userRepo, cfg)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		ErrorHandler: http.ErrorHandler,
	})

	// Setup routes
	http.SetupRoutes(app, userUsecase, cfg)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Social Network API! - JakeOnyx")
	})

	// Start server
	port := cfg.App.Port
	log.Printf("Starting %s in %s mode on port %s", cfg.App.Name, cfg.App.Env, cfg.App.Port)
	log.Fatal(app.Listen(":" + port))

}
