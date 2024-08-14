package loginwithmagiclinkusecase

import (
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (uc *loginWithMagicLinkUseCase) Execute(req request.LoginWithMagicLinkRequest) (string, string, errors.AppErr) {
	logger.Info("Init LoginWithMagicLink UseCase")

	magicLink := uc.mlRepo.FindByValue(req.Code)
	if magicLink == nil {
		return "", "", errors.NewAppErr("Code not found", http.StatusNotFound)
	}

	if time.Now().After(magicLink.GetExpiresAt()) {
		return "", "", errors.NewAppErr("Code has expired", http.StatusUnauthorized)
	}

	account := uc.repo.FindById(magicLink.GetAccountID())
	if account == nil {
		return "", "", errors.NewAppErr("Account not found", http.StatusNotFound)
	}

	accessToken, refreshToken, tokenErr := uc.atProvider.GenerateAuthTokens(account.GetID())
	if tokenErr != nil {
		return "", "", tokenErr
	}

	var lastCountry, lastCity string
	if account.GetLastLoginCountry() != nil {
		lastCountry = *account.GetLastLoginCountry()
	}
	if account.GetLastLoginCity() != nil {
		lastCity = *account.GetLastLoginCity()
	}

	country, city, geoErr := uc.glProvider.GetInfo(req.IP)
	if geoErr != nil {
		logger.Error("Error trying to retrive geolocation data", geoErr)
		return "", "", errors.NewInternalServerErr()
	}

	if lastCountry != "" && (lastCountry != country || lastCity != city) {
		msg := "Sua conta foi acessada em outra localização. caso não tenha sido você recomendamos que altere sua senha. obrigado pela atenção!"
		uc.emailProvider.SendMail(account.GetEmail(), "login suspeito.", msg)
	}

	deviceDetails := utils.GetDeviceDetails(req.UserAgent)
	now := time.Now()

	device := uc.devicesRepo.FindByIpAndAccountId(req.IP, account.GetID())

	if device == nil {
		device = entities.NewDevice(
			account.GetID(),
			deviceDetails.Name,
			req.UserAgent,
			deviceDetails.Platform,
			req.IP,
			now,
		)

		err := uc.devicesRepo.Create(device)
		if err != nil {
			logger.Error("Error trying to create device", err)
			return "", "", errors.NewInternalServerErr()
		}
	} else {
		device.LoginNow()
		uc.devicesRepo.Update(device)
	}

	account.LoginNow()
	account.SetLastLoginCountry(country)
	account.SetLastLoginCity(city)
	account.SetLastLoginIp(req.IP)

	err := uc.repo.Update(account)
	if err != nil {
		logger.Error("Error trying to update account data.", err)
		return "", "", errors.NewInternalServerErr()
	}

	err = uc.mlRepo.Delete(magicLink.GetID())
	if err != nil {
		logger.Error("Error while delete magic link", err)
		return "", "", errors.NewInternalServerErr()
	}

	return accessToken, refreshToken, nil
}
