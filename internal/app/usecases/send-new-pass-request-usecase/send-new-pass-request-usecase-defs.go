package sendnewpassrequestusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type sendNewPassRequestUseCase struct {
	repo          contracts.AccountsRepository
	vcRepo        contracts.VerificationCodesRepository
	emailProvider contracts.EmailProvider
}

func NewSendNewPassRequestUseCase(
	repo contracts.AccountsRepository,
	vcRepo contracts.VerificationCodesRepository,
	emailProvider contracts.EmailProvider,
) usecases.SendNewPassRequestUseCase {
	return &sendNewPassRequestUseCase{
		repo:          repo,
		vcRepo:        vcRepo,
		emailProvider: emailProvider,
	}
}
