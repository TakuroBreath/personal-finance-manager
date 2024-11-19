package service

import (
	"fmt"
	"github.com/TakuroBreath/personal-finance-manager/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService struct {
	secretKey   []byte
	tokenExpiry time.Duration
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{
		secretKey:   []byte(secretKey),
		tokenExpiry: 24 * time.Hour,
	}
}

func (s *JWTService) GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(s.tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *JWTService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Проверка времени истечения токена
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, fmt.Errorf("token expired")
		}
	}

	return claims, nil
}
