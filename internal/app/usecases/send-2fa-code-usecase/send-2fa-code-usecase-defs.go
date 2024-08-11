package send2facodeusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type send2faCodeUsecase struct {
	repo                  contracts.AccountsRepository
	vcRepo                contracts.VerificationCodesRepository
	twoFactorAuthProvider contracts.TwoFactorAuthProvider
}

func NewSend2faCodeUseCase(
	repo contracts.AccountsRepository,
	vcRepo contracts.VerificationCodesRepository,
	twoFactorAuthProvider contracts.TwoFactorAuthProvider,
) usecases.Send2faCodeUseCase {
	return &send2faCodeUsecase{
		repo:                  repo,
		vcRepo:                vcRepo,
		twoFactorAuthProvider: twoFactorAuthProvider,
	}
}
