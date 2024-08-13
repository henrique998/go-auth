package providers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type authTokensProvider struct {
	rtRepo contracts.RefreshTokensRepository
}

func NewAuthTokensProvider(rtRepo contracts.RefreshTokensRepository) contracts.AuthTokensProvider {
	return &authTokensProvider{rtRepo: rtRepo}
}

func (p *authTokensProvider) GenerateAuthTokens(accountId string) (string, string, errors.AppErr) {
	tokenExpiresAt := time.Now().Add(15 * time.Minute)
	accessToken, tokenErr := utils.GenerateJWTToken(accountId, tokenExpiresAt, os.Getenv("JWT_SECRET"))
	if tokenErr != nil {
		logger.Error("Error trying to generate access token token", tokenErr)
		return "", "", errors.NewInternalServerErr()
	}

	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24 * 30)
	refreshToken, tokenErr := utils.GenerateJWTToken(accountId, refreshTokenExpiresAt, os.Getenv("JWT_SECRET"))
	if tokenErr != nil {
		logger.Error("Error trying to generate refresh token", tokenErr)
		return "", "", errors.NewInternalServerErr()
	}

	rt := entities.NewRefreshToken(refreshToken, accountId, refreshTokenExpiresAt)

	p.rtRepo.Create(rt)

	return accessToken, refreshToken, nil
}

func (p *authTokensProvider) ValidateJWTToken(tokenString string) (string, errors.AppErr) {
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
