package createaccountcontroller

import (
	"github.com/henrique998/go-auth-2/internal/app/usecases"
	"github.com/henrique998/go-auth-2/internal/infra/controllers"
)

type createAccountController struct {
	uc usecases.CreateAccountUsecase
}

func NewCreateAccountController(uc usecases.CreateAccountUsecase) controllers.CreateAccountController {
	return &createAccountController{uc: uc}
}
