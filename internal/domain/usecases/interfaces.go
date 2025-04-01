package usecases

import (
	"context"

	"github.com/huuloc2026/go-social/internal/domain/entities"
)

type UserUseCase interface {
	GetUserByID(ctx context.Context, id uint) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type AuthUseCase interface {
	Register(ctx context.Context, req *entities.RegisterRequest) (*entities.AuthResponse, error)
	Login(ctx context.Context, req *entities.LoginRequest) (*entities.AuthResponse, error)
	ValidateToken(token string) (uint, error)
}
