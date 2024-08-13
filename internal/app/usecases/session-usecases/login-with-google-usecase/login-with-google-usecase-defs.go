package loginwithgoogleusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type loginWithGoogleUseCase struct {
	repo          contracts.AccountsRepository
	emailProvider contracts.EmailProvider
	atProvider    contracts.AuthTokensProvider
	devicesRepo   contracts.DevicesRepository
	glProvider    contracts.GeoLocationProvider
}

func NewLoginWithGoogleUseCase(
	repo contracts.AccountsRepository,
	emailProvider contracts.EmailProvider,
	atProvider contracts.AuthTokensProvider,
	devicesRepo contracts.DevicesRepository,
	glProvider contracts.GeoLocationProvider,
) usecases.LoginWithGoogleUseCase {
	return &loginWithGoogleUseCase{
		repo:          repo,
		emailProvider: emailProvider,
		atProvider:    atProvider,
		devicesRepo:   devicesRepo,
		glProvider:    glProvider,
	}
}
