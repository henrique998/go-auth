package usecases

import (
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
)

type LoginWithCredentialsUseCase interface {
	Execute(req request.LoginWithCredentialsRequest) (string, string, errors.AppErr)
}

type LoginWithGoogleUseCase interface {
	Execute(data request.LoginWithGoogleRequest) (string, string, errors.AppErr)
}

type LoginWithMagicLinkUseCase interface {
	Execute(req request.LoginWithMagicLinkRequest) (string, string, errors.AppErr)
}
