package usecases

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/entities"
)

type userUseCase struct {
	userRepo UserRepository
}

func NewUserUseCase(userRepo UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uc *userUseCase) GetUserByID(ctx context.Context, id uint) (*entities.User, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	return user, nil
}

func (uc *userUseCase) UpdateUser(ctx context.Context, user *entities.User) error {
	existingUser, err := uc.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	// Update only allowed fields
	existingUser.Name = user.Name
	existingUser.Email = user.Email

	if err := uc.userRepo.Update(ctx, existingUser); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update user")
	}

	return nil
}

func (uc *userUseCase) DeleteUser(ctx context.Context, id uint) error {
	if err := uc.userRepo.Delete(ctx, id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete user")
	}
	return nil
}
