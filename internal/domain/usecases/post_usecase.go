package usecases

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"github.com/huuloc2026/go-social/internal/interfaces/http/repositories"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, userID uint, content string, images []string, videos []string) (*entities.Post, error)
	GetPostByID(ctx context.Context, id uint) (*entities.Post, error)
	GetAllPosts(ctx context.Context, offset, limit int) ([]entities.Post, error)
	UpdatePost(ctx context.Context, id uint, content string, images []string, videos []string) (*entities.Post, error)
	DeletePost(ctx context.Context, id uint) error
}

type postUseCase struct {
	postRepo repositories.PostRepository
}

func NewPostUseCase(postRepo repositories.PostRepository) PostUseCase {
	return &postUseCase{postRepo}
}

func (uc *postUseCase) CreatePost(ctx context.Context, userID uint, content string, images []string, videos []string) (*entities.Post, error) {
	post := &entities.Post{
		UserID:  userID,
		Content: content,
		Images:  images,
		Videos:  videos,
	}
	return uc.postRepo.Create(ctx, post)
}

func (uc *postUseCase) GetPostByID(ctx context.Context, id uint) (*entities.Post, error) {
	return uc.postRepo.GetByID(ctx, id)
}

func (uc *postUseCase) GetAllPosts(ctx context.Context, offset, limit int) ([]entities.Post, error) {
	return uc.postRepo.GetAll(ctx, offset, limit)
}

func (uc *postUseCase) UpdatePost(ctx context.Context, id uint, content string, images []string, videos []string) (*entities.Post, error) {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Post not found")
	}

	post.Content = content
	post.Images = images
	post.Videos = videos

	return uc.postRepo.Update(ctx, post)
}

func (uc *postUseCase) DeletePost(ctx context.Context, id uint) error {
	return uc.postRepo.Delete(ctx, id)
}
