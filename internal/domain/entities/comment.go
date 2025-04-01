package entities

import "time"

// Comment struct represents a Comment entity for a post
type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	PostID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
