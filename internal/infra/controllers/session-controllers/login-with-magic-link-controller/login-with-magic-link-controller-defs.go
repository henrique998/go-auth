package loginwithmagiclinkcontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type loginWithMagicLinkController struct {
	usecase usecases.LoginWithMagicLinkUseCase
}

func NewLoginWithMagicLinkController(
	usecase usecases.LoginWithMagicLinkUseCase,
) controllers.Controller {
	return &loginWithMagicLinkController{usecase: usecase}
}
