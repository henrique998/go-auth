package updatepasscontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type updatePassController struct {
	usecase usecases.UpdatePassUseCase
}

func NewUpdatePassController(usecase usecases.UpdatePassUseCase) controllers.Controller {
	return &updatePassController{usecase: usecase}
}
