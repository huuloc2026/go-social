package main

import (
	"log"

	"github.com/huuloc2026/go-social/infrastructure/database"
	"github.com/huuloc2026/go-social/internal/interfaces/http"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/config"
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

	// Initialize Redis cache
	//redisCache := cache.NewRedisCache(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)

	// Initialize RabbitMQ
	//rabbitMQ := queue.NewRabbitMQ(cfg.RabbitMQ.Host, cfg.RabbitMQ.Port, cfg.RabbitMQ.User, cfg.RabbitMQ.Password)
	//defer rabbitMQ.Close()

	// Initialize repositories
	userRepo := domain.NewUserRepository(db.DB)
	//postRepo := domain.NewPostRepository(db.DB)
	// friendshipRepo := domain.NewFriendshipRepository(db.DB)
	// commentRepo := domain.NewCommentRepository(db.DB)
	// likeRepo := domain.NewLikeRepository(db.DB)
	// notificationRepo := domain.NewNotificationRepository(db.DB)

	// Initialize services
	//notificationService := domain.NewNotificationService(notificationRepo, rabbitMQ)

	// Initialize use cases
	userUsecase := usecases.NewUserUsecase(userRepo, cfg)
	//postUsecase := usecases.NewPostUsecase(postRepo, notificationService)
	// friendshipUsecase := usecases.NewFriendshipUsecase(friendshipRepo, userRepo, notificationService)
	// interactionUsecase := usecases.NewInteractionUsecase(commentRepo, likeRepo, postRepo, notificationService)
	// feedUsecase := usecases.NewFeedUsecase(postRepo, friendshipRepo, redisCache)
	// notificationUsecase := usecases.NewNotificationUsecase(notificationRepo)

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
