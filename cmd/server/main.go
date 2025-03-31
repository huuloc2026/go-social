package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/adapters/db"
	"github.com/huuloc2026/go-social/internal/adapters/http"
	"github.com/huuloc2026/go-social/internal/config"
	"github.com/huuloc2026/go-social/internal/core/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	dbConn, err := db.NewPostgresDB(cfg.PostgresDB)
	if err != nil {
		panic(err)
	}

	// Initialize services
	userService := services.NewUserService(dbConn)

	// Setup Fiber app
	app := fiber.New()

	// Register HTTP handlers
	http.NewUserHandler(app, userService)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Social Network API! - JakeOnyx")
	})
	// Start server
	port := strconv.Itoa(cfg.AppConfig.Port)
	log.Fatal(app.Listen(":" + port))
}
