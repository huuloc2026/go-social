package repositories

import (
	"context"

	"github.com/huuloc2026/go-social/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByID(ctx context.Context, id uint) (*entities.User, error)
	FindAllWithPagination(ctx context.Context, offset, limit int) ([]*entities.User, error)
	CountAll(ctx context.Context) (int, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uint) error
}
