package verify2facodeusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type verify2faCodeUseCase struct {
	repo   contracts.AccountsRepository
	vcRepo contracts.VerificationCodesRepository
}

func NewVerify2faCodeUseCase(
	repo contracts.AccountsRepository,
	vcRepo contracts.VerificationCodesRepository,
) usecases.Verify2faCodeUseCase {
	return &verify2faCodeUseCase{
		repo:   repo,
		vcRepo: vcRepo,
	}
}
