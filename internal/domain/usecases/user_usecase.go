package usecases

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"github.com/huuloc2026/go-social/internal/domain/repositories"
)

type userUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uc *userUseCase) GetUserByID(ctx context.Context, id uint) (*entities.User, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	return user, nil
}

func (uc *userUseCase) GetAllUsers(ctx context.Context, page, limit int) ([]*entities.User, int, error) {
	// Validate pagination parameters
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	// Calculate offset for pagination
	offset := (page - 1) * limit

	// Fetch users with pagination
	users, err := uc.userRepo.FindAllWithPagination(ctx, offset, limit)
	if err != nil {
		return nil, 0, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch users")
	}

	// Get the total count of users (for pagination info)
	totalCount, err := uc.userRepo.CountAll(ctx)
	if err != nil {
		return nil, 0, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch total user count")
	}

	// Return the users and total count for pagination
	return users, totalCount, nil
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
