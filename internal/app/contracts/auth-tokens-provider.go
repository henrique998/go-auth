package contracts

import "github.com/henrique998/go-auth/internal/app/errors"

type AuthTokensProvider interface {
	GenerateAuthTokens(accountId string) (string, string, errors.AppErr)
	ValidateJWTToken(token string) (string, errors.AppErr)
}
