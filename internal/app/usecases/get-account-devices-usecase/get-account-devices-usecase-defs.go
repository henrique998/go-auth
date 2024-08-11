package getaccountdevicesusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type getAccountDevicesUsecase struct {
	repo        contracts.AccountsRepository
	devicesRepo contracts.DevicesRepository
}

func NewGetAccountDevicesUseCase(repo contracts.AccountsRepository, devicesRepo contracts.DevicesRepository) usecases.GetAccountDevicesUsecase {
	return &getAccountDevicesUsecase{
		repo:        repo,
		devicesRepo: devicesRepo,
	}
}
