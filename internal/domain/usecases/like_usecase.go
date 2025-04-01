package usecases

import (
	"github.com/huuloc2026/go-social/internal/domain/errors"
	"github.com/huuloc2026/go-social/internal/interfaces/http/repositories"
)

type LikeUseCase interface {
	LikePost(userID, postID uint) error
	UnlikePost(userID, postID uint) error
	GetLikeCount(postID uint) (uint, error)
}

type likeUseCaseImpl struct {
	likeRepo repositories.LikeRepository
}

func NewLikeUseCase(likeRepo repositories.LikeRepository) LikeUseCase {
	return &likeUseCaseImpl{likeRepo: likeRepo}
}

// Like một post
func (uc *likeUseCaseImpl) LikePost(userID, postID uint) error {
	if err := uc.likeRepo.LikePost(userID, postID); err != nil {
		return errors.ErrBadRequest
	}
	return nil
}

// Unlike một post
func (uc *likeUseCaseImpl) UnlikePost(userID, postID uint) error {
	if err := uc.likeRepo.UnlikePost(userID, postID); err != nil {
		return errors.ErrNotFound
	}
	return nil
}

// Lấy tổng số likes của một post
func (uc *likeUseCaseImpl) GetLikeCount(postID uint) (uint, error) {
	return uc.likeRepo.CountLikes(postID)
}
