package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/huuloc2026/go-social/internal/domain/entities"
)

var (
	jwtSecret         []byte
	jwtExpiration     time.Duration
	refreshExpiration time.Duration
	initOnce          sync.Once // Đảm bảo InitJWT chỉ được gọi 1 lần
)

type Claims struct {
	UserID uint          `json:"user_id"`
	Role   entities.Role `json:"role"`
	jwt.RegisteredClaims
}

// InitJWT khởi tạo JWT với secret key và thời gian hết hạn
func InitJWT(secret string, exp time.Duration, refreshExp time.Duration) {
	initOnce.Do(func() {
		jwtSecret = []byte(secret)
		jwtExpiration = exp
		refreshExpiration = refreshExp
	})
}

// GenerateJWT tạo access token với thông tin user
func GenerateJWT(userID uint, role entities.Role) (string, error) {
	if len(jwtSecret) == 0 {
		return "", errors.New("JWT secret not initialized")
	}

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// GenerateRefreshToken tạo refresh token ngẫu nhiên an toàn
func GenerateRefreshToken() (string, error) {
	token := uuid.New().String()
	return token, nil
}

// ParseToken kiểm tra và giải mã JWT
func ParseToken(tokenString string) (*Claims, error) {
	if len(jwtSecret) == 0 {
		return nil, errors.New("JWT secret not initialized")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GenerateRandomToken tạo token ngẫu nhiên với độ dài tùy chỉnh
func GenerateRandomToken(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, length)

	for i := range token {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		token[i] = charset[randomIndex.Int64()]
	}

	return string(token), nil
}

// ValidateJWT xác thực token và trả về UserID nếu hợp lệ
func ValidateJWT(tokenString string) (uint, error) {
	if len(jwtSecret) == 0 {
		return 0, errors.New("JWT secret not initialized")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid token")
}
