package utils

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateResetPasswordToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(API_SECRET))
}

func ParseResetToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exceptions.NewCustomError(http.StatusBadRequest, "invalid or expired token")
		}
		return []byte(API_SECRET), nil
	})

	if err != nil {
		return nil, exceptions.NewCustomError(http.StatusBadRequest, "invalid or expired token")
	}

	if !token.Valid {
		return nil, exceptions.NewCustomError(http.StatusBadRequest, "invalid or expired token")
	}

	return claims, nil
}
