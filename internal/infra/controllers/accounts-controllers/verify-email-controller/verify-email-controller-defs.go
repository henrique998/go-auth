package verifyemailcontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type verifyEmailController struct {
	usecase usecases.VerifyEmailUseCase
}

func NewVerifyEmailController(
	usecase usecases.VerifyEmailUseCase,
) controllers.Controller {
	return &verifyEmailController{usecase: usecase}
}
