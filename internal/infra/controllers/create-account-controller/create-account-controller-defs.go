package createaccountcontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type createAccountController struct {
	uc usecases.CreateAccountUsecase
}

func NewCreateAccountController(uc usecases.CreateAccountUsecase) controllers.Controller {
	return &createAccountController{uc: uc}
}
