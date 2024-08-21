package verify2facodecontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type verify2FACodeController struct {
	usecase usecases.Verify2faCodeUseCase
}

func NewVerify2FACodeController(
	usecase usecases.Verify2faCodeUseCase,
) controllers.Controller {
	return &verify2FACodeController{usecase: usecase}
}
