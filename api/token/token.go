package token

import (
	"errors"
	"time"
	"weather/pkg/model"

	"github.com/golang-jwt/jwt/v5"
)

type Claim struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string) (*model.GetTokenResp, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := Claim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte("Hamidjon0424"))
	if err != nil {
		return nil, err
	}
	return &model.GetTokenResp{
		Token:     accessToken,
		ExpiresAt: expirationTime.Format("2006-01-02 15:04:05.999999999-07:00"),
	}, nil
}

func ExtractClaimToken(stringToken string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(stringToken, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("Hamidjon0424"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ValidToken(tokenString string) (bool, error) {
	_, err := ExtractClaimToken(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}
