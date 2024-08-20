package sendnewpassrequestcontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type sendNewPassRequestController struct {
	usecase usecases.SendNewPassRequestUseCase
}

func NewSendNewPassRequestController(
	usecase usecases.SendNewPassRequestUseCase,
) controllers.Controller {
	return &sendNewPassRequestController{usecase: usecase}
}
