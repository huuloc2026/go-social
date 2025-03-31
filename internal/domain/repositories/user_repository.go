package repositories

import (
	"github.com/huuloc2026/go-social/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error
	return &user, err
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	return repo.DB.Save(user).Error
}

func (repo *UserRepository) DeleteUser(id uint) error {
	return repo.DB.Delete(&models.User{}, id).Error
}
