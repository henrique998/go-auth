package updatepassusecase

import (
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (uc *updatePassUsecase) Execute(req request.NewPassRequest) errors.AppErr {
	logger.Info("Init UpdatePass UseCase")

	code := uc.vcRepo.FindByValue(req.Code)

	if code == nil {
		return errors.NewAppErr("code not found", http.StatusNotFound)
	}

	account := uc.repo.FindById(code.GetAccountID())

	now := time.Now()

	if now.After(code.GetExpiresAt()) {
		return errors.NewAppErr("code has expired", http.StatusUnauthorized)
	}

	if req.NewPass != req.NewPassConfirmation {
		return errors.NewAppErr("new password and confirmation must be equals", http.StatusBadRequest)
	}

	if len(req.NewPass) < 6 {
		return errors.NewAppErr("new password must contain 6 or more characters", http.StatusBadRequest)
	}

	if account.GetPass() != nil && utils.ComparePassword(req.NewPass, *account.GetPass()) {
		return errors.NewAppErr("new password cannot be the same as the previous one", http.StatusBadRequest)
	}

	newPassHash, err := utils.HashPass(req.NewPass)
	if err != nil {
		logger.Error("Error trying to hash new password", err)
		return errors.NewInternalServerErr()
	}

	account.SetPass(newPassHash)
	account.Touch()

	uc.repo.Update(account)
	uc.vcRepo.Delete(code.GetID())

	return nil
}
