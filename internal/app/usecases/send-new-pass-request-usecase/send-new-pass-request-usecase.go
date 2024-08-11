package sendnewpassrequestusecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	appErr "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (uc *sendNewPassRequestUseCase) Execute(email string) appErr.AppErr {
	logger.Info("Init SendNewPassRequest UseCase")

	account := uc.repo.FindByEmail(email)

	if account == nil {
		return appErr.NewAppErr("account does not exists", http.StatusNotFound)
	}

	code, err := utils.GenerateCode(10)
	if err != nil {
		logger.Error("Error trying to generate code!", err)
		return appErr.NewAppErr("internal server error", http.StatusInternalServerError)
	}

	message := fmt.Sprintf("Your verification code is: %s. Please use this code to complete your password reset process. The code will expire in 10 minutes.", code)
	expiresAt := time.Now().Add(10 * time.Minute)

	verificationCode := entities.NewVerificationCode(code, account.GetID(), expiresAt)

	uc.vcRepo.Create(verificationCode)
	uc.emailProvider.SendMail(email, "Password reset", message)

	return nil
}
