package usecases

import "github.com/huuloc2026/go-social/internal/domain"

type PostUsecase interface {
	CreatePost(userID uint, content string) (*domain.Post, error)
	GetPostByID(id uint) (*domain.Post, error)
	GetPostsByUserID(userID uint, limit, offset int) ([]*domain.Post, error)
	UpdatePost(userID, postID uint, content string) (*domain.Post, error)
	DeletePost(userID, postID uint) error
}

type postUsecase struct {
	postRepo domain.PostRepository
}

func NewPostUsecase(postRepo domain.PostRepository) PostUsecase {
	return &postUsecase{postRepo: postRepo}
}

func (uc *postUsecase) CreatePost(userID uint, content string) (*domain.Post, error) {
	post := &domain.Post{
		UserID:  userID,
		Content: content,
	}

	if err := uc.postRepo.Create(post); err != nil {
		return nil, err
	}

	return uc.postRepo.FindByID(post.ID)
}

func (uc *postUsecase) GetPostByID(id uint) (*domain.Post, error) {
	return uc.postRepo.FindByID(id)
}

func (uc *postUsecase) GetPostsByUserID(userID uint, limit, offset int) ([]*domain.Post, error) {
	return uc.postRepo.FindByUserID(userID, limit, offset)
}

func (uc *postUsecase) UpdatePost(userID, postID uint, content string) (*domain.Post, error) {
	post, err := uc.postRepo.FindByID(postID)
	if err != nil {
		return nil, err
	}

	if post.UserID != userID {
		return nil, domain.ErrUnauthorized
	}

	post.Content = content
	if err := uc.postRepo.Update(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *postUsecase) DeletePost(userID, postID uint) error {
	post, err := uc.postRepo.FindByID(postID)
	if err != nil {
		return err
	}

	if post.UserID != userID {
		return domain.ErrUnauthorized
	}

	return uc.postRepo.Delete(postID)
}
