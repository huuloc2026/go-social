package repositories

import (
	"errors"

	"github.com/huuloc2026/go-social/internal/domain/entities"
	domainErrors "github.com/huuloc2026/go-social/internal/domain/errors"
	"gorm.io/gorm"
)

type LikeRepository interface {
	LikePost(userID, postID uint) error
	UnlikePost(userID, postID uint) error
	CountLikes(postID uint) (uint, error)
}

type likeRepositoryImpl struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepositoryImpl{db: db}
}

// Like một post
func (r *likeRepositoryImpl) LikePost(userID, postID uint) error {
	// Kiểm tra xem user đã like post này chưa
	var existingLike entities.Like
	if err := r.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&existingLike).Error; err == nil {
		return errors.New("user already liked this post")
	}

	// Nếu chưa like, thêm bản ghi vào bảng likes
	like := entities.Like{UserID: userID, PostID: postID}
	if err := r.db.Create(&like).Error; err != nil {
		return err
	}

	// Tăng likes_count của bài post
	return r.db.Model(&entities.Post{}).Where("id = ?", postID).Update("likes_count", gorm.Expr("likes_count + 1")).Error
}

// Unlike một post
// Unlike một post
func (r *likeRepositoryImpl) UnlikePost(userID, postID uint) error {
	// Kiểm tra xem user đã like post này chưa
	var like entities.Like
	if err := r.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&like).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return domainErrors.ErrNotFound
		}

		return err
	}

	// Nếu đã like, xóa bản ghi trong bảng likes
	if err := r.db.Delete(&like).Error; err != nil {
		return err
	}

	// Giảm likes_count của bài post
	return r.db.Model(&entities.Post{}).Where("id = ?", postID).Update("likes_count", gorm.Expr("likes_count - 1")).Error
}

func (r *likeRepositoryImpl) CountLikes(postID uint) (uint, error) {
	var count int64
	err := r.db.Model(&entities.Like{}).Where("post_id = ?", postID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return uint(count), nil
}
