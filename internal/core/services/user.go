package services

import (
	"github.com/huuloc2026/go-social/internal/core/domain"
	"github.com/huuloc2026/go-social/internal/core/ports"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(username, email string) (*domain.User, error) {
	user := domain.NewUser(0, username, email)
	err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUser(id int) (*domain.User, error) {
	return s.repo.FindByID(id)
}
