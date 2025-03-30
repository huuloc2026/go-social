package ports

import "github.com/huuloc2026/go-social/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindByID(id int) (*domain.User, error)
}

type UserService interface {
	CreateUser(username, email string) (*domain.User, error)
	GetUser(id int) (*domain.User, error)
}
