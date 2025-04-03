package usecases

import (
	"context"

	"github.com/huuloc2026/go-social/internal/domain/entities"
)

type UserUseCase interface {
	GetUserByID(ctx context.Context, id uint) (*entities.User, error)
	GetAllUsers(ctx context.Context, page, limit int) ([]*entities.User, int, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type AuthUseCase interface {
	Register(ctx context.Context, req *entities.RegisterRequest) (*entities.AuthResponse, error)
	Login(ctx context.Context, req *entities.LoginRequest) (*entities.AuthResponse, error)

	ValidateToken(token string) (uint, error)
	//Who Are You
	WhoAreYou(ctx context.Context, userId uint) error
	// Refresh an access token using a refresh token
	RefreshToken(refreshToken string) (*entities.AuthResponse, error)

	// Reset a user's password using a reset token
	ResetPassword(ctx context.Context, token string, newPassword string) error

	// Change a user's password (usually requires the user to provide current password)
	ChangePassword(ctx context.Context, userID uint, oldPassword string, newPassword string) error

	// Logout a user (usually invalidates the session or refresh token)
	Logout(ctx context.Context, userID uint) error

	// Verify an email address for a user (could be used during registration)
	VerifyEmail(ctx context.Context, userID uint, verificationToken string) error
}
