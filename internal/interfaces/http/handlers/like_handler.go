package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/utils"
)

type LikeHandler struct {
	likeUseCase usecases.LikeUseCase
}

func NewLikeHandler(likeUseCase usecases.LikeUseCase) *LikeHandler {
	return &LikeHandler{likeUseCase: likeUseCase}
}

// Like một post
func (h *LikeHandler) LikePost(ctx *fiber.Ctx) error {
	userID, _ := utils.ExtractUserID(ctx)

	postID, err := strconv.Atoi(ctx.Params("post_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid post ID")
	}

	if err := h.likeUseCase.LikePost(uint(userID), uint(postID)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Post liked successfully"})
}

// Unlike một post
func (h *LikeHandler) UnlikePost(ctx *fiber.Ctx) error {
	userID, _ := utils.ExtractUserID(ctx)

	postID, err := strconv.Atoi(ctx.Params("post_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid post ID")
	}

	if err := h.likeUseCase.UnlikePost(uint(userID), uint(postID)); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Post unliked successfully"})
}

// Lấy tổng số likes của một post
func (h *LikeHandler) GetLikeCount(ctx *fiber.Ctx) error {
	postID, err := strconv.Atoi(ctx.Params("post_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid post ID")
	}

	count, err := h.likeUseCase.GetLikeCount(uint(postID))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch like count")
	}

	return ctx.JSON(fiber.Map{"likes": count})
}
