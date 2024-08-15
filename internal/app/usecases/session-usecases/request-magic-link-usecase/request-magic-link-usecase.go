package requestmagiclinkusecase

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (uc *requestMagicLinkUseCase) Execute(email string) errors.AppErr {
	logger.Info("Init RequestMagicLink UseCase")

	account := uc.repo.FindByEmail(email)
	if account == nil {
		return errors.NewAppErr("account not found", http.StatusNotFound)
	}

	code, err := utils.GenerateCode(32)
	if err != nil {
		logger.Error("Error trying to generate magic link code", err)
		return errors.NewAppErr("internal server error!", http.StatusInternalServerError)
	}
	expiresAt := time.Now().Add(15 * time.Minute)

	magicLink := entities.NewMagicLink(account.GetID(), code, expiresAt)

	err = uc.mlRepo.Create(magicLink)
	if err != nil {
		logger.Error("Error trying to create magic link", err)
		return errors.NewInternalServerErr()
	}

	magicLinkURL := fmt.Sprintf("%s/session/login/magic-link?code=%s", os.Getenv("BASE_URL"), code)
	body := fmt.Sprintf("Click the link to login: %s\nThis link will expire in 15 minutes.", magicLinkURL)

	err = uc.emailProvider.SendMail(account.GetEmail(), "Auth with magic link!", body)
	if err != nil {
		logger.Error("Error trying to send email with magic link", err)
		return errors.NewInternalServerErr()
	}

	return nil
}
