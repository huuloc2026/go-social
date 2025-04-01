package usecases

import (
	"context"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"github.com/huuloc2026/go-social/internal/domain/repositories"
	"github.com/huuloc2026/go-social/internal/utils"
)

type authUseCase struct {
	userRepo repositories.UserRepository
	// mailer        mail.NewNodeMailer
	// cache         cache.RedisCache
	refreshExpiry time.Duration
}

func NewAuthUseCase(userRepo repositories.UserRepository, refreshExpiry time.Duration) AuthUseCase {
	return &authUseCase{
		userRepo:      userRepo,
		refreshExpiry: refreshExpiry,
	}
}

func (uc *authUseCase) Register(ctx context.Context, req *entities.RegisterRequest) (*entities.AuthResponse, error) {
	// Validate input
	if err := utils.ValidateRequest(req); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Check if user exists
	existingUser, _ := uc.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Email already in use")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	// Create user
	user := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     entities.RoleUser,
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	// Generate verification token
	// verificationToken, err := utils.GenerateRefreshToken()
	// if err != nil {
	// 	return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate verification token")
	// }

	// Save verification token
	// if err := uc.tokenRepo.Create(ctx, &entities.Token{
	// 	UserID:    user.ID,
	// 	Token:     verificationToken,
	// 	Type:      "verification",
	// 	ExpiresAt: time.Now().Add(24 * time.Hour),
	// }); err != nil {
	// 	return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to save verification token")
	// }

	// Send verification email (async)
	//go uc.mailer.SendVerificationEmail(user.Email, verificationToken)

	// Generate tokens
	accessToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate access token")
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate refresh token")
	}

	// Save refresh token
	// if err := uc.tokenRepo.Create(ctx, &entities.Token{
	// 	UserID:    user.ID,
	// 	Token:     refreshToken,
	// 	Type:      "refresh",
	// 	ExpiresAt: time.Now().Add(uc.refreshExpiry),
	// }); err != nil {
	// 	return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to save refresh token")
	// }

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update last login")
	}

	return &entities.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (uc *authUseCase) Login(ctx context.Context, req *entities.LoginRequest) (*entities.AuthResponse, error) {
	// Validate request
	if err := utils.ValidateRequest(req); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Find user by email
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	// Verify password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	// Generate access token
	accessToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate access token")
	}

	// Generate refresh token
	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate refresh token")
	}

	// // Save refresh token to database
	// err = uc.tokenRepo.Create(ctx, &entities.Token{
	// 	UserID:    user.ID,
	// 	Token:     refreshToken,
	// 	Type:      "refresh",
	// 	ExpiresAt: time.Now().Add(uc.refreshExpiry),
	// })
	// if err != nil {
	// 	return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to save refresh token")
	// }

	// Update last login timestamp
	now := time.Now()
	user.LastLogin = &now
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update last login")
	}

	// Return response
	return &entities.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *authUseCase) ValidateToken(token string) (uint, error) {
	return utils.ValidateJWT(token)
}
