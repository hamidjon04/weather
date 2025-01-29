package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization header ni olish
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Tokenni yaroqliligini tekshirish
		claims, err := ExtractClaimToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		// Tokenning yaroqlilik muddatini tekshirish
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			c.Abort()
			return
		}

		// Claims ni kontekstga qo'shish
		c.Set("claims", claims)
		c.Next()
	}
}

type Claim struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

// ExtractClaimToken - Tokenni claims qismini `Claim` structiga o'tkazish
func ExtractClaimToken(tokenString string) (*Claim, error) {
	// Tokenni parse qilish
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		// Tokenni imzolash kalitini qaytarish (masalan, secret key)
		return []byte("your_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}

	// Claims ni olish
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
