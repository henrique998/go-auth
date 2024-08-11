package verify2facodeusecase

import (
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

func (uc *verify2faCodeUseCase) Execute(req request.Verify2faCodeRequest) errors.AppErr {
	logger.Info("Init Verify2faCode UseCase")

	verificationCode := uc.vcRepo.FindByValue(req.Code)

	if verificationCode == nil {
		return errors.NewAppErr("verification code not found", http.StatusNotFound)
	}

	if verificationCode.GetAccountID() != req.AccountId {
		return errors.NewAppErr("unauthorized action", http.StatusUnauthorized)
	}

	now := time.Now()

	if verificationCode.GetExpiresAt().Before(now) {
		return errors.NewAppErr("verification code has expired", http.StatusUnauthorized)
	}

	account := uc.repo.FindById(req.AccountId)

	if account.GetIs2FaEnabled() {
		return errors.NewAppErr("Two factor authentication already carried out", http.StatusUnauthorized)
	}

	account.Enable2FA()
	account.Touch()

	uc.repo.Update(account)

	uc.vcRepo.Delete(verificationCode.GetID())

	return nil
}
