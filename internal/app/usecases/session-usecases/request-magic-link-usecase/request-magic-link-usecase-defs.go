package requestmagiclinkusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type requestMagicLinkUseCase struct {
	repo          contracts.AccountsRepository
	mlRepo        contracts.MagicLinksRepository
	emailProvider contracts.EmailProvider
}

func NewRequestMagicLinkUseCase(
	repo contracts.AccountsRepository,
	mlRepo contracts.MagicLinksRepository,
	emailProvider contracts.EmailProvider,
) usecases.RequestMagicLinkUseCase {
	return &requestMagicLinkUseCase{
		repo:          repo,
		mlRepo:        mlRepo,
		emailProvider: emailProvider,
	}
}
