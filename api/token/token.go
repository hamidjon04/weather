package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte("Hamidjon0424"))
	if err != nil{
		return "", err
	}
	return accessToken, nil
}
