package loginwithmagiclinkusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type loginWithMagicLinkUseCase struct {
	repo          contracts.AccountsRepository
	mlRepo        contracts.MagicLinksRepository
	devicesRepo   contracts.DevicesRepository
	atProvider    contracts.AuthTokensProvider
	emailProvider contracts.EmailProvider
	glProvider    contracts.GeoLocationProvider
}

func NewLoginWithMagicLinkUseCase(
	repo contracts.AccountsRepository,
	mlRepo contracts.MagicLinksRepository,
	devicesRepo contracts.DevicesRepository,
	atProvider contracts.AuthTokensProvider,
	emailProvider contracts.EmailProvider,
	glProvider contracts.GeoLocationProvider,
) usecases.LoginWithMagicLinkUseCase {
	return &loginWithMagicLinkUseCase{
		repo:          repo,
		mlRepo:        mlRepo,
		devicesRepo:   devicesRepo,
		atProvider:    atProvider,
		emailProvider: emailProvider,
		glProvider:    glProvider,
	}
}
