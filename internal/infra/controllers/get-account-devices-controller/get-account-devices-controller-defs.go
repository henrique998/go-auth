package getaccountdevicescontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type getAccountDevicesController struct {
	usecase usecases.GetAccountDevicesUsecase
}

func NewGetAccountDevicesController(
	usecase usecases.GetAccountDevicesUsecase,
) controllers.Controller {
	return &getAccountDevicesController{usecase: usecase}
}
