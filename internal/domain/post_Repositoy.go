package domain

import "gorm.io/gorm"

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}

func (r *PostRepositoryImpl) Create(post *Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepositoryImpl) FindByID(id uint) (*Post, error) {
	var post Post
	err := r.db.Preload("User").Preload("Comments").Preload("Likes").First(&post, id).Error
	return &post, err
}

func (r *PostRepositoryImpl) FindByUserID(userID uint, limit, offset int) ([]*Post, error) {
	var posts []*Post
	err := r.db.Preload("User").
		Preload("Comments").
		Preload("Likes").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error
	return posts, err
}

func (r *PostRepositoryImpl) Update(post *Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&Post{}, id).Error
}

func (r *PostRepositoryImpl) GetFeed(userIDs []uint, limit, offset int) ([]*Post, error) {
	var posts []*Post
	err := r.db.Preload("User").
		Preload("Comments").
		Preload("Likes").
		Where("user_id IN ?", userIDs).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error
	return posts, err
}
