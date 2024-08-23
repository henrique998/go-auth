package refreshtokencontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type refreshTokenController struct {
	usecase usecases.RefreshTokenUseCase
}

func NewRefreshTokenController(
	usecase usecases.RefreshTokenUseCase,
) controllers.Controller {
	return &refreshTokenController{usecase: usecase}
}
