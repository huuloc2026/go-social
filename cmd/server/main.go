package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/config"

	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/infrastructure/database"
	"github.com/huuloc2026/go-social/internal/interfaces/http"
	"github.com/huuloc2026/go-social/internal/interfaces/http/repositories"
	"github.com/huuloc2026/go-social/internal/utils"
)

func main() {
	//////////////////////////////////////////////////
	//                                              //
	//              Load configuration              //
	//                                              //
	//////////////////////////////////////////////////

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	//////////////////////////////////////////////////
	//                                              //
	//               Initialize JWT                 //
	//                                              //
	//////////////////////////////////////////////////

	utils.InitJWT(cfg.JWTSecret, cfg.JWTExpiration, cfg.RefreshExpiration)

	//////////////////////////////////////////////////
	//                                              //
	//             Initialize database              //
	//                                              //
	//////////////////////////////////////////////////

	db, err := database.NewPostgresDB(&cfg)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	//////////////////////////////////////////////////
	//                                              //
	//                  Service Mail                //
	//                                              //
	//////////////////////////////////////////////////

	//mailer := mail.NewNodeMailer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPassword)

	//////////////////////////////////////////////////
	//                                              //
	//                Service Cache                 //
	//                                              //
	//////////////////////////////////////////////////
	//cache := cache.NewRedisCache(cfg.RedisURL)

	refreshExpiry := time.Duration(cfg.RefreshExpiration) * time.Second

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)
	likeRepo := repositories.NewLikeRepository(db)

	// Initialize use cases
	userUseCase := usecases.NewUserUseCase(userRepo)
	postUseCase := usecases.NewPostUseCase(postRepo)
	likeUseCase := usecases.NewLikeUseCase(likeRepo)
	authUseCase := usecases.NewAuthUseCase(userRepo, refreshExpiry)

	// Create Fiber app

	app := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://localhost:3000",      // Allow frontend to access backend (replace with your frontend URL)
	// 	AllowMethods:     "GET,POST,PUT,DELETE",        // Allowed methods
	// 	AllowHeaders:     "Content-Type,Authorization", // Allowed headers
	// 	AllowCredentials: true,                         // Allow credentials (cookies, HTTP authentication)
	// }))

	// Setup routes
	http.SetupRoutes(app, authUseCase, userUseCase, postUseCase, likeUseCase)

	//////////////////////////////////////////////////
	//                                              //
	//                 Start server                 //
	//                                              //
	//////////////////////////////////////////////////
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Healthcheck OK")
	})
	port := ":" + cfg.Port
	if err := app.Listen(port); err != nil {
		log.Fatal("Failed to start server:", err)
		os.Exit(1)
	}
}
