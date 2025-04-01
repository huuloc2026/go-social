package entities

import "time"

// Comment struct represents a Comment entity for a post
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null" json:"postId"`
	UserID    uint      `gorm:"not null" json:"userId"`
	Content   string    `gorm:"type:text;not null" json:"Content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
