package usecases

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"github.com/huuloc2026/go-social/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userRepo UserRepository
}

func NewAuthUseCase(userRepo UserRepository) AuthUseCase {
	return &authUseCase{userRepo: userRepo}
}

func (uc *authUseCase) Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error) {
	// Check if user already exists
	existingUser, _ := uc.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Email already in use")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	// Create user
	user := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	return user, nil
}

func (uc *authUseCase) Login(ctx context.Context, req *entities.LoginRequest) (string, error) {
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token")
	}

	return token, nil
}

func (uc *authUseCase) ValidateToken(token string) (uint, error) {
	return utils.ValidateJWT(token)
}
