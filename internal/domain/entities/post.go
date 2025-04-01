package entities

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content"`
	Image     string    `json:"image"` // Chỉ giữ lại 1 string cho ảnh
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Likes     int       `json:"likes"`
}
