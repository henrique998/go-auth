package send2facodecontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type send2FACodeController struct {
	usecase usecases.Send2faCodeUseCase
}

func NewSend2FACodeController(
	usecase usecases.Send2faCodeUseCase,
) controllers.Controller {
	return &send2FACodeController{usecase: usecase}
}
