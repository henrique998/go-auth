package usecases

import (
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
)

type CreateAccountUsecase interface {
	Execute(req request.CreateAccountRequest) errors.AppErr
}

type GetAccountDevicesUsecase interface {
	Execute(accountId string) ([]entities.Device, errors.AppErr)
}
