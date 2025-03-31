package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type NotificationType string

const (
	NotificationTypeFriendRequest NotificationType = "friendship_request"
	NotificationTypePostLike      NotificationType = "post_like"
	NotificationTypeNewComment    NotificationType = "new_comment"
)

type Notification struct {
	ID         uint             `gorm:"primaryKey" json:"id"`
	UserID     uint             `gorm:"not null" json:"user_id"`
	FromUserID uint             `gorm:"not null" json:"from_user_id"`
	Type       NotificationType `gorm:"type:varchar(50);not null" json:"type"`
	Content    JSON             `gorm:"type:jsonb" json:"content"`
	Read       bool             `gorm:"default:false" json:"read"`
	CreatedAt  time.Time        `gorm:"autoCreateTime" json:"created_at"`

	User     *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	FromUser *User `gorm:"foreignKey:FromUserID" json:"from_user,omitempty"`
}

type NotificationService interface {
	Send(toUserID, fromUserID uint, notificationType NotificationType, content map[string]interface{}) error
	MarkAsRead(userID, notificationID uint) error
	GetNotifications(userID uint, limit, offset int) ([]*Notification, error)
	GetUnreadCount(userID uint) (int64, error)
}

// JSON type for handling JSONB in PostgreSQL
type JSON map[string]interface{}

func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &j)
}
