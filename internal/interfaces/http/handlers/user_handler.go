package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/utils"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (c *UserHandler) GetUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid user ID")
	}

	user, err := c.userUseCase.GetUserByID(ctx.Context(), uint(id))
	if err != nil {
		return utils.ErrorResponse(ctx, err.(*fiber.Error).Code, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)    // Default to page 1 if not provided
	limit := c.QueryInt("limit", 10) // Default to 10 items per page if not provided

	users, total, err := h.userUseCase.GetAllUsers(c.Context(), page, limit)
	if err != nil {
		return err
	}

	// Calculate total pages
	totalPages := (total + limit - 1) / limit

	return c.JSON(fiber.Map{
		"data":         users,
		"total":        total,
		"total_pages":  totalPages,
		"current_page": page,
	})
}

func (c *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid user ID")
	}

	var user entities.User
	if err := ctx.BodyParser(&user); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	user.ID = uint(id)
	if err := c.userUseCase.UpdateUser(ctx.Context(), &user); err != nil {
		return utils.ErrorResponse(ctx, err.(*fiber.Error).Code, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, fiber.Map{"message": "User updated successfully"})
}

func (c *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid user ID")
	}

	if err := c.userUseCase.DeleteUser(ctx.Context(), uint(id)); err != nil {
		return utils.ErrorResponse(ctx, err.(*fiber.Error).Code, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, fiber.Map{"message": "User deleted successfully"})
}
