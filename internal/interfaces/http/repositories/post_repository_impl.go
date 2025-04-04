package repositories

import (
	"context"

	"gorm.io/gorm"

	"github.com/huuloc2026/go-social/internal/domain/entities"
	"github.com/huuloc2026/go-social/internal/domain/repositories"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repositories.PostRepository {
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
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAll(ctx context.Context, offset, limit int) ([]entities.Post, error) {
	var posts []entities.Post
	if err := r.db.Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
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
