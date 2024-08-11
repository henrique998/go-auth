package verifyemailaccountusecase

import (
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

func (uc *verifyEmailUseCase) Execute(token string) errors.AppErr {
	logger.Info("Init VerifyEmail UseCase")

	verificationToken := uc.vcRepo.FindByValue(token)

	if time.Now().After(verificationToken.GetExpiresAt()) {
		return errors.NewAppErr("verification code has expired", http.StatusUnauthorized)
	}

	account := uc.repo.FindById(verificationToken.GetAccountID())

	if account.GetIsEmailVerified() {
		return errors.NewAppErr("the email has already been verified", http.StatusUnauthorized)
	}

	account.VerifyEmail()
	account.Touch()

	uc.repo.Update(account)

	uc.vcRepo.Delete(verificationToken.GetID())

	return nil
}
