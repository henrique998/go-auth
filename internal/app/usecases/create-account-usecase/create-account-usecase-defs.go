package createaccountusecase

import (
	"github.com/henrique998/go-auth-2/internal/app/repositories"
	"github.com/henrique998/go-auth-2/internal/app/usecases"
)

type createAccountUsecase struct {
	repo repositories.AccountsRepository
}

func NewCreateAccountUseCase(repo repositories.AccountsRepository) usecases.CreateAccountUsecase {
	return &createAccountUsecase{repo: repo}
}
