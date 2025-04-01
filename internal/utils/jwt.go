package utils

import (
	"errors"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/huuloc2026/go-social/internal/domain/entities"
)

var (
	jwtSecret         []byte
	jwtExpiration     time.Duration
	refreshExpiration time.Duration
)

type Claims struct {
	UserID uint          `json:"user_id"`
	Role   entities.Role `json:"role"`
	jwt.RegisteredClaims
}

func InitJWT(secret string, exp time.Duration, refreshExp time.Duration) {
	jwtSecret = []byte(secret)
	jwtExpiration = exp
	refreshExpiration = refreshExp
}

func GenerateJWT(userID uint, role entities.Role) (string, error) {
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

func GenerateRefreshToken() (string, error) {
	refreshToken := uuid.New().String()
	return refreshToken, nil
}

func ParseToken(tokenString string) (*Claims, error) {
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
func GenerateRandomToken(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, length)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token), nil
}
