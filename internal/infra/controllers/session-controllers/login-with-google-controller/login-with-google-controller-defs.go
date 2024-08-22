package loginwithgooglecontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type loginWithGoogleController struct {
	usecase usecases.LoginWithGoogleUseCase
}

func NewLoginWithGoogleController(
	usecase usecases.LoginWithGoogleUseCase,
) controllers.Controller {
	return &loginWithGoogleController{usecase: usecase}
}
