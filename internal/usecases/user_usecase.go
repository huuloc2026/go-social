package usecases

import (
	"errors"

	"github.com/huuloc2026/go-social/internal/domain/models"
	"github.com/huuloc2026/go-social/internal/domain/repositories"
	"github.com/huuloc2026/go-social/pkg/utils"
)

type UserUsecase struct {
	UserRepo *repositories.UserRepository
}

func NewUserUsecase(userRepo *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

// Đăng ký tài khoản mới
func (uc *UserUsecase) Register(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return uc.UserRepo.CreateUser(user)
}

// Đăng nhập và tạo JWT token
func (uc *UserUsecase) Login(email, password string) (string, error) {
	user, err := uc.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Lấy thông tin user
func (uc *UserUsecase) GetUserByID(id uint) (*models.User, error) {
	return uc.UserRepo.GetUserByID(id)
}

// Cập nhật thông tin user
func (uc *UserUsecase) UpdateUser(user *models.User) error {
	return uc.UserRepo.UpdateUser(user)
}

// Xóa user
func (uc *UserUsecase) DeleteUser(id uint) error {
	return uc.UserRepo.DeleteUser(id)
}
