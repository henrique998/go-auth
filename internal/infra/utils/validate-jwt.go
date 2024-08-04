package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/henrique998/go-auth-2/internal/app/errors"
)

func ValidateJWTToken(tokenString string) (string, errors.AppErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", errors.NewAppErr("invalid token", http.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.NewAppErr("invalid token", http.StatusUnauthorized)
	}

	accountId := claims["sub"].(string)
	exp := int64(claims["exp"].(float64))

	if exp < time.Now().Unix() {
		return "", errors.NewAppErr("token expired", http.StatusUnauthorized)
	}

	return accountId, nil
}
