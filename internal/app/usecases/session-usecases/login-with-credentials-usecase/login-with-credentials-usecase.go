package loginwithcredentialsusecase

import (
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (uc *loginWithCredentialsUseCase) Execute(req request.LoginWithCredentialsRequest) (string, string, errors.AppErr) {
	logger.Info("Init LoginWithCredentials UseCase")

	account := uc.repo.FindByEmail(req.Email)
	if account == nil {
		la := entities.NewLoginAttempt(req.Email, req.IP, req.UserAgent, false)

		err := uc.laRepository.Create(la)
		if err != nil {
			logger.Error("Error trying to create login attempt record.", err)
		}

		return "", "", errors.NewAppErr("email or password incorrect!", http.StatusBadRequest)
	}

	if account.GetPass() == nil {
		la := entities.NewLoginAttempt(req.Email, req.IP, req.UserAgent, false)

		err := uc.laRepository.Create(la)
		if err != nil {
			logger.Error("Error trying to create login attempt record.", err)
		}

		return "", "", errors.NewAppErr("Login method not allowed!", http.StatusUnauthorized)
	}

	passwordMatch := utils.ComparePassword(req.Pass, *account.GetPass())
	if !passwordMatch {
		la := entities.NewLoginAttempt(req.Email, req.IP, req.UserAgent, false)

		err := uc.laRepository.Create(la)
		if err != nil {
			logger.Error("Error trying to create login attempt record.", err)
		}

		return "", "", errors.NewAppErr("email or password incorrect!", 400)
	}

	if !account.GetIsEmailVerified() {
		la := entities.NewLoginAttempt(req.Email, req.IP, req.UserAgent, false)

		err := uc.laRepository.Create(la)
		if err != nil {
			logger.Error("Error trying to create login attempt record.", err)
		}

		return "", "", errors.NewAppErr("Only verified accounts can log in!", http.StatusUnauthorized)
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

	country, city, err := uc.glProvider.GetInfo(req.IP)
	if err != nil {
		logger.Error("Error trying to retrive geolocation data", err)
		return "", "", errors.NewInternalServerErr()
	}

	if lastCountry != "" && (lastCountry != country || lastCity != city) {
		msg := "Sua conta foi acessada em outra localização. caso não tenha sido você recomendamos que altere sua senha. obrigado pela atenção!"
		uc.emailProvider.SendMail(req.Email, "login suspeito", msg)
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

		err = uc.devicesRepo.Create(device)
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

	err = uc.repo.Update(account)
	if err != nil {
		logger.Error("Error trying to update account data.", err)
		return "", "", errors.NewInternalServerErr()
	}

	la := entities.NewLoginAttempt(account.GetEmail(), req.IP, req.UserAgent, true)

	err = uc.laRepository.Create(la)
	if err != nil {
		logger.Error("Error trying to create login attempt record.", err)
		return accessToken, refreshToken, nil
	}

	return accessToken, refreshToken, nil
}
