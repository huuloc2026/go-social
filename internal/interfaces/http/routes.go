package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/interfaces/http/handlers"
	"github.com/huuloc2026/go-social/internal/interfaces/http/middlewares"
)

func SetupRoutes(app *fiber.App,
	authUseCase usecases.AuthUseCase,
	userUseCase usecases.UserUseCase,
	postUseCase usecases.PostUseCase,
	likeUseCase usecases.LikeUseCase,
) {
	authHandler := handlers.NewAuthController(authUseCase)
	userHandler := handlers.NewUserController(userUseCase)
	postHanlder := handlers.NewPostController(postUseCase)
	likeHandler := handlers.NewLikeHandler(likeUseCase)
	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// User routes
	user := app.Group("/users")
	user.Use(middlewares.AuthMiddleware(authUseCase))
	user.Get("/", userHandler.GetAllUsers)
	user.Get("/:id", userHandler.GetUser)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)

	// Post Routes
	post := app.Group("/posts")
	post.Use(middlewares.AuthMiddleware(authUseCase))
	post.Post("/", postHanlder.CreatePost)
	post.Get("/", postHanlder.GetAllPosts)
	post.Get("/:id", postHanlder.GetPostByID)
	post.Put("/:id", postHanlder.UpdatePost)
	post.Delete("/:id", postHanlder.DeletePost)
	//Like Routes
	like := app.Group("/like")
	like.Use(middlewares.AuthMiddleware(authUseCase))
	app.Post("/:post_id/like", likeHandler.LikePost)
	app.Delete("/:post_id/like", likeHandler.UnlikePost)
	app.Get("/:post_id/likes", likeHandler.GetLikeCount)
}
