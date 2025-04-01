package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/utils"
)

type UserController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(userUseCase usecases.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
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

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
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

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid user ID")
	}

	if err := c.userUseCase.DeleteUser(ctx.Context(), uint(id)); err != nil {
		return utils.ErrorResponse(ctx, err.(*fiber.Error).Code, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, fiber.Map{"message": "User deleted successfully"})
}
