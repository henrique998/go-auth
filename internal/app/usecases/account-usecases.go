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

type Send2faCodeUseCase interface {
	Execute(accountId string) errors.AppErr
}

type SendNewPassRequestUseCase interface {
	Execute(email string) errors.AppErr
}

type UpdatePassUseCase interface {
	Execute(req request.NewPassRequest) errors.AppErr
}
