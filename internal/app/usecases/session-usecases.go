package usecases

import (
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
)

type LoginWithCredentialsUseCase interface {
	Execute(req request.LoginWithCredentialsRequest) (string, string, errors.AppErr)
}
