package entities

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	gorm.Model
	Name      string     `gorm:"size:100;not null" json:"name"`
	Email     string     `gorm:"size:100;unique;not null" json:"email"`
	Password  string     `gorm:"size:255;not null" json:"-"`
	Role      Role       `gorm:"type:varchar(20);default:'user'" json:"role"`
	Verified  bool       `gorm:"default:false" json:"verified"`
	LastLogin *time.Time `json:"last_login,omitempty"`
}

type Token struct {
	gorm.Model
	UserID    uint      `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"size:512;not null" json:"token"`
	Type      string    `gorm:"size:50;not null" json:"type"` // refresh, verification, password_reset
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Used      bool      `gorm:"default:false" json:"used"`
}

// Request/Response DTOs
type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AuthResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPassword struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
