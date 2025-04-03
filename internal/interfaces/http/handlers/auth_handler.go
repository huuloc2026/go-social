package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"github.com/huuloc2026/go-social/internal/domain/usecases"
	"github.com/huuloc2026/go-social/internal/utils"
)

type AuthController struct {
	authUseCase usecases.AuthUseCase
}

func NewAuthController(authUseCase usecases.AuthUseCase) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var req entities.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := utils.ValidateRequest(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	user, err := c.authUseCase.Register(ctx.Context(), &req)
	if err != nil {
		return utils.ErrorResponse(ctx, err.(*fiber.Error).Code, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusCreated, user)
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req entities.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := utils.ValidateRequest(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	token, err := c.authUseCase.Login(ctx.Context(), &req)
	if err != nil {
		return utils.ErrorResponse(ctx, err.(*fiber.Error).Code, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, token)
}

func (c *AuthController) Me(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(uint)
	if !ok {
		return utils.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid user ID type")
	}
	// Nếu cần lấy dữ liệu từ body/query/params, xử lý ở đây

	//user := c.authUseCase.WhoAreYou(ctx.Context(), userID)
	// if err != nil {
	// 	return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	// }
	// if err != nil {
	// 	if fiberErr, ok := err.(*fiber.Error); ok {
	// 		return utils.ErrorResponse(ctx, fiberErr.Code, fiberErr.Error())
	// 	}
	// 	return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	// }

	return utils.SuccessResponse(ctx, fiber.StatusOK, userID)
}
