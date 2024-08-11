package send2facodeusecase

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (uc *send2faCodeUsecase) Execute(accountId string) appError.AppErr {
	logger.Info("Init Send2FACode UseCase")

	account := uc.repo.FindById(accountId)

	if account == nil {
		return appError.NewAppErr("account does not exists", http.StatusNotFound)
	}

	if account.GetPhone() == nil {
		return appError.NewAppErr("account must have an phone number to complete 2fa proccess", http.StatusUnauthorized)
	}

	code, err := utils.GenerateCode(10)
	if err != nil {
		logger.Error("Error while generate 2fa code.", err)
		return appError.NewInternalServerErr()
	}

	message := fmt.Sprintf("Your verification code is: %s. Please use this code to complete your 2FA process. The code will expire in 10 minutes.", code)
	expiresAt := time.Now().Add(10 * time.Minute)
	fromNumber := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	verificationCode := entities.NewVerificationCode(code, accountId, expiresAt)

	uc.vcRepo.Create(verificationCode)
	err = uc.twoFactorAuthProvider.Send(fromNumber, *account.GetPhone(), message)
	if err != nil {
		logger.Error("Error trying to send 2fa code!", err)
		return appError.NewInternalServerErr()
	}

	return nil
}
