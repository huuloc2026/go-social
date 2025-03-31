package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/huuloc2026/go-social/internal/config"
	"github.com/huuloc2026/go-social/internal/interfaces/adapters/database"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup Fiber app
	app := fiber.New()

	//Connect DB
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Social Network API! - JakeOnyx")
	})
	// Start server
	port := cfg.App.Port
	log.Printf("Starting %s in %s mode on port %s", cfg.App.Name, cfg.App.Env, cfg.App.Port)
	log.Fatal(app.Listen(":" + port))
}
