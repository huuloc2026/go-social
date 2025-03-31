package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/config"
	"github.com/huuloc2026/go-social/internal/interfaces/http"
	"github.com/spf13/viper"
)

func main() {
	// Load config
	config.LoadConfig()

	app := fiber.New()

	// Setup routes
	http.SetupRoutes(app)

	port := viper.GetString("app.port")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
