package usecases

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/huuloc2026/go-social/internal/domain"
	"github.com/huuloc2026/go-social/pkg/utils"
)

// internal/usecases/auth_usecase.go
type AuthUseCase struct {
	userRepo  domain.UserRepository
	jwtSecret string
	jwtExpiry time.Duration
}

func (uc *AuthUseCase) Login(email, password string) (string, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(uc.jwtExpiry).Unix(),
	})

	return token.SignedString([]byte(uc.jwtSecret))
}
