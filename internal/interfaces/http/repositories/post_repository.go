package repositories

import (
	"context"

	"github.com/huuloc2026/go-social/internal/domain/entities"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(ctx context.Context, post *entities.Post) (*entities.Post, error)
	GetByID(ctx context.Context, id uint) (*entities.Post, error)
	GetAll(ctx context.Context, offset, limit int) ([]entities.Post, error)
	Update(ctx context.Context, post *entities.Post) (*entities.Post, error)
	Delete(ctx context.Context, id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(ctx context.Context, post *entities.Post) (*entities.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) GetByID(ctx context.Context, id uint) (*entities.Post, error) {
	var post entities.Post
	if err := r.db.Preload("Comments").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAll(ctx context.Context, offset, limit int) ([]entities.Post, error) {
	var posts []entities.Post
	if err := r.db.Offset(offset).Limit(limit).Preload("Comments").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) Update(ctx context.Context, post *entities.Post) (*entities.Post, error) {
	if err := r.db.Save(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.Delete(&entities.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
