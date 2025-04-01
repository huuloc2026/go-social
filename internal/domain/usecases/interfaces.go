package usecases

import (
	"context"

	"github.com/huuloc2026/go-social/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByID(ctx context.Context, id uint) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uint) error
}

type UserUseCase interface {
	GetUserByID(ctx context.Context, id uint) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type AuthUseCase interface {
	Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error)
	Login(ctx context.Context, req *entities.LoginRequest) (string, error)
	ValidateToken(token string) (uint, error)
}
