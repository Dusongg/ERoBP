package common

import (
	"ERoBP/modle"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

func GenerateJWT(userID uint, username, role string, expTime time.Duration) (string, error) {
	expirationTime := time.Now().Add(expTime)
	claims := &modle.Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Id:        uuid.New().String(), // 确保每个 Token 的唯一性
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(modle.JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
