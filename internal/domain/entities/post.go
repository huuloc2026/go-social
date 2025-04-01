package entities

import (
	"time"
)

type Post struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	Content    string    `json:"content"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LikesCount uint      `json:"likes_count"`
}
