package createaccountusecase

import (
	"fmt"

	"github.com/henrique998/go-auth-2/internal/app/errors"
	"github.com/henrique998/go-auth-2/internal/app/request"
)

func (uc *createAccountUsecase) Execute(req request.CreateAccountRequest) errors.AppErr {
	fmt.Println(req)

	return nil
}
