package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
)

type PostHandler struct {
	postUseCase usecases.PostUseCase
}

func NewPostController(postUseCase usecases.PostUseCase) *PostHandler {
	return &PostHandler{postUseCase: postUseCase}
}

func (c *PostHandler) CreatePost(ctx *fiber.Ctx) error {
	var request struct {
		Content string   `json:"content"`
		Images  []string `json:"images"`
		Videos  []string `json:"videos"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	userID := 1 // For now, you can pass the userID from the session or JWT token
	post, err := c.postUseCase.CreatePost(ctx.Context(), uint(userID), request.Content, request.Images, request.Videos)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(post)
}

func (c *PostHandler) GetPostByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid post ID")
	}

	post, err := c.postUseCase.GetPostByID(ctx.Context(), uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}

func (c *PostHandler) GetAllPosts(ctx *fiber.Ctx) error {
	offset, _ := strconv.Atoi(ctx.Query("offset", "0"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	posts, err := c.postUseCase.GetAllPosts(ctx.Context(), offset, limit)
	if err != nil {
		return err
	}

	return ctx.JSON(posts)
}

func (c *PostHandler) UpdatePost(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid post ID")
	}

	var request struct {
		Content string   `json:"content"`
		Images  []string `json:"images"`
		Videos  []string `json:"videos"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	post, err := c.postUseCase.UpdatePost(ctx.Context(), uint(id), request.Content, request.Images, request.Videos)
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}

func (c *PostHandler) DeletePost(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid post ID")
	}

	if err := c.postUseCase.DeletePost(ctx.Context(), uint(id)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
