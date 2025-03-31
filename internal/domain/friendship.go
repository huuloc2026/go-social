package domain

import (
	"time"
)

type FriendshipStatus string

const (
	FriendshipPending  FriendshipStatus = "pending"
	FriendshipAccepted FriendshipStatus = "accepted"
	FriendshipRejected FriendshipStatus = "rejected"
)

type Friendship struct {
	ID        uint             `gorm:"primaryKey" json:"id"`
	UserID    uint             `gorm:"not null" json:"user_id"`
	FriendID  uint             `gorm:"not null" json:"friend_id"`
	Status    FriendshipStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CreatedAt time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time        `gorm:"autoUpdateTime" json:"updated_at"`

	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Friend *User `gorm:"foreignKey:FriendID" json:"friend,omitempty"`
}

type FriendshipRepository interface {
	Create(friendship *Friendship) error
	FindByID(id uint) (*Friendship, error)
	FindByUserID(userID uint) ([]*Friendship, error)
	FindByUsers(userID, friendID uint) (*Friendship, error)
	Update(friendship *Friendship) error
	Delete(id uint) error
	GetFriends(userID uint) ([]*User, error)
	GetFriendRequests(userID uint) ([]*Friendship, error)
}
