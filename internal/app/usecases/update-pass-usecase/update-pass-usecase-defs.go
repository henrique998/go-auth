package updatepassusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type updatePassUsecase struct {
	repo   contracts.AccountsRepository
	vcRepo contracts.VerificationCodesRepository
}

func NewUpdatePassUseCase(
	repo contracts.AccountsRepository,
	vcRepo contracts.VerificationCodesRepository,
) usecases.UpdatePassUseCase {
	return &updatePassUsecase{
		repo:   repo,
		vcRepo: vcRepo,
	}
}
