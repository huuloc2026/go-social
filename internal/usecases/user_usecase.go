package usecases

import (
	"errors"
	"time"

	"github.com/huuloc2026/go-social/internal/config"
	"github.com/huuloc2026/go-social/internal/domain"
	"github.com/huuloc2026/go-social/pkg/utils"

	"github.com/golang-jwt/jwt/v4"
)

type UserUsecase interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error)
	GetUserByID(id uint) (*domain.User, error)
}

type userUsecase struct {
	userRepo domain.UserRepository
	config   *config.Config
}

func NewUserUsecase(userRepo domain.UserRepository, cfg *config.Config) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		config:   cfg,
	}
}

func (uc *userUsecase) Register(user *domain.User) error {
	// Check if email already exists
	_, err := uc.userRepo.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	// Check if username already exists
	_, err = uc.userRepo.FindByUsername(user.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	return uc.userRepo.Create(user)
}

func (uc *userUsecase) Login(email, password string) (string, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	})

	return token.SignedString([]byte(uc.config.App.JWTSecret))
}

func (uc *userUsecase) GetUserByID(id uint) (*domain.User, error) {
	return uc.userRepo.FindByID(id)
}
