package repositories

import (
	"context"

	"github.com/huuloc2026/go-social/internal/domain/entities"
)

type PostRepository interface {
	Create(ctx context.Context, post *entities.Post) (*entities.Post, error)
	GetByID(ctx context.Context, id uint) (*entities.Post, error)
	GetAll(ctx context.Context, offset, limit int) ([]entities.Post, error)
	Update(ctx context.Context, post *entities.Post) (*entities.Post, error)
	Delete(ctx context.Context, id uint) error
}
