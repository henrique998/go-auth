package refreshtokenusecase

import (
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

func (uc *refreshTokenUseCase) Execute(refreshToken string) (string, string, errors.AppErr) {
	logger.Info("Init RefreshToken UseCase")

	accountId, err := uc.atRepo.ValidateJWTToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	newAccessToken, newRefreshToken, err := uc.atRepo.GenerateAuthTokens(accountId)
	if err != nil {
		return "", "", err
	}

	uc.repo.Delete(refreshToken)

	return newAccessToken, newRefreshToken, nil
}
