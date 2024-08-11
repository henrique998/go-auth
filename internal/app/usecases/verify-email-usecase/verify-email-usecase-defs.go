package verifyemailaccountusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type verifyEmailUseCase struct {
	repo   contracts.AccountsRepository
	vcRepo contracts.VerificationCodesRepository
}

func NewVerifyEmailUseCase(
	repo contracts.AccountsRepository,
	vcRepo contracts.VerificationCodesRepository,
) usecases.VerifyEmailUseCase {
	return &verifyEmailUseCase{
		repo:   repo,
		vcRepo: vcRepo,
	}
}
