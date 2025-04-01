package entities

import "time"

// Post struct represents a Post entity in the system
type Post struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"` // Foreign key to User
	Content   string    `gorm:"type:text;not null"`
	Images    []string  `gorm:"type:text[]"`
	Videos    []string  `gorm:"type:text[]"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Likes     uint      `gorm:"default:0"`
	Comments  []Comment // Relationship with Comment entity
}
