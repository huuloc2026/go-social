package domain

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Associations
	User     *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
	Likes    []Like    `gorm:"foreignKey:PostID" json:"likes,omitempty"`
}

type PostRepository interface {
	Create(post *Post) error
	FindByID(id uint) (*Post, error)
	FindByUserID(userID uint, limit, offset int) ([]*Post, error)
	Update(post *Post) error
	Delete(id uint) error
	GetFeed(userIDs []uint, limit, offset int) ([]*Post, error)
}
