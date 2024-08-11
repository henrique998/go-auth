package getaccountdevicesusecase

import (
	"net/http"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

func (uc *getAccountDevicesUsecase) Execute(accountId string) ([]entities.Device, errors.AppErr) {
	logger.Info("Init GetAccountDevices UseCase")

	account := uc.repo.FindById(accountId)

	if account == nil {
		return nil, errors.NewAppErr("account does not exists", http.StatusNotFound)
	}

	devices := uc.devicesRepo.FindManyByAccountId(accountId)

	return devices, nil
}
