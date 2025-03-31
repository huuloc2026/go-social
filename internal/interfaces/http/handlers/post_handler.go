package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/usecases"
	"github.com/huuloc2026/go-social/pkg/utils"
)

type PostHandler struct {
	postUsecase usecases.PostUsecase
}

func NewPostHandler(postUsecase usecases.PostUsecase) *PostHandler {
	return &PostHandler{postUsecase: postUsecase}
}

type CreatePostRequest struct {
	Content string `json:"content" validate:"required,min=1,max=1000"`
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	post, err := h.postUsecase.CreatePost(userID, req.Content)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create post",
		})
	}

	return c.Status(http.StatusCreated).JSON(post)
}

func (h *PostHandler) GetPost(c *fiber.Ctx) error {
	postID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}

	post, err := h.postUsecase.GetPostByID(uint(postID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	return c.JSON(post)
}

// Add similar handlers for UpdatePost, DeletePost, GetUserPosts
