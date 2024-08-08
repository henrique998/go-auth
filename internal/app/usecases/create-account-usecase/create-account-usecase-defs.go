package createaccountusecase

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/usecases"
)

type createAccountUsecase struct {
	repo          contracts.AccountsRepository
	vcRepo        contracts.VerificationCodesRepository
	emailProvider contracts.EmailProvider
}

func NewCreateAccountUseCase(repo contracts.AccountsRepository, vcRepo contracts.VerificationCodesRepository, emailProvider contracts.EmailProvider) usecases.CreateAccountUsecase {
	return &createAccountUsecase{
		repo:          repo,
		vcRepo:        vcRepo,
		emailProvider: emailProvider,
	}
}
