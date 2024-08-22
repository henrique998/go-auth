package loginwithcredentialscontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type loginWithCredentialsController struct {
	usecase usecases.LoginWithCredentialsUseCase
}

func NewLoginWithCredentialsController(
	usecase usecases.LoginWithCredentialsUseCase,
) controllers.Controller {
	return &loginWithCredentialsController{usecase: usecase}
}
