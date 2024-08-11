package loginwithcredentialsusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type loginWithCredentialsUseCase struct {
	repo          contracts.AccountsRepository
	devicesRepo   contracts.DevicesRepository
	laRepository  contracts.LoginAttemptsRepository
	emailProvider contracts.EmailProvider
	atProvider    contracts.AuthTokensProvider
	glProvider    contracts.GeoLocationProvider
}

func NewLoginWithCredentialsUseCase(
	repo contracts.AccountsRepository,
	devicesRepo contracts.DevicesRepository,
	laRepository contracts.LoginAttemptsRepository,
	emailProvider contracts.EmailProvider,
	atProvider contracts.AuthTokensProvider,
	glProvider contracts.GeoLocationProvider,
) usecases.LoginWithCredentialsUseCase {
	return &loginWithCredentialsUseCase{
		repo:          repo,
		devicesRepo:   devicesRepo,
		laRepository:  laRepository,
		emailProvider: emailProvider,
		atProvider:    atProvider,
		glProvider:    glProvider,
	}
}
