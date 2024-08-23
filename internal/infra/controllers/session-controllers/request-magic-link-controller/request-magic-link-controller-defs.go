package requestmagiclinkcontroller

import (
	"github.com/henrique998/go-auth/internal/app/usecases"
	"github.com/henrique998/go-auth/internal/infra/controllers"
)

type requestMagicLinkController struct {
	usecase usecases.RequestMagicLinkUseCase
}

func NewRequestMagicLinkController(
	usecase usecases.RequestMagicLinkUseCase,
) controllers.Controller {
	return &requestMagicLinkController{usecase: usecase}
}
