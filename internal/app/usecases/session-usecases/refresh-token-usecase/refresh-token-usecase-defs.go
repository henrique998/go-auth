package refreshtokenusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type refreshTokenUseCase struct {
	repo   contracts.RefreshTokensRepository
	atRepo contracts.AuthTokensProvider
}

func NewRefreshTokenUseCase(
	repo contracts.RefreshTokensRepository,
	atRepo contracts.AuthTokensProvider,
) usecases.RefreshTokenUseCase {
	return &refreshTokenUseCase{
		repo:   repo,
		atRepo: atRepo,
	}
}
