package usecases

import (
	"github.com/henrique998/go-auth-2/internal/app/errors"
	"github.com/henrique998/go-auth-2/internal/app/request"
)

type CreateAccountUsecase interface {
	Execute(req request.CreateAccountRequest) errors.AppErr
}
