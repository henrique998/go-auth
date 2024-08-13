package loginwithgoogleusecase

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/providers"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

var userInfoUrl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
var personFieldsUrl = "https://people.googleapis.com/v1/people/me?personFields=phoneNumbers"

func (uc *loginWithGoogleUseCase) Execute(data request.LoginWithGoogleRequest) (string, string, errors.AppErr) {
	logger.Info("Init LoginWithGoogle UseCase")

	accessToken, err := utils.GetGoogleAccessToken(data.Code)
	if err != nil {
		logger.Error("Error trying while get google oauth access token", err)
		return "", "", errors.NewInternalServerErr()
	}

	res, err := http.Get(userInfoUrl + accessToken)
	if err != nil {
		logger.Error("Error trying while get user info using access token", err)
		return "", "", errors.NewInternalServerErr()
	}
	defer res.Body.Close()

	userInfo := make(map[string]any)
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		logger.Error("failed to decode user info:", err)
		return "", "", errors.NewInternalServerErr()
	}

	req, err := http.NewRequest("GET", personFieldsUrl, nil)
	if err != nil {
		logger.Error("failed to generate http request:", err)
		return "", "", errors.NewInternalServerErr()
	}

	gcProvider := providers.GoogleHTTPClientProvider{
		Token: accessToken,
	}

	res, err = gcProvider.Do(req)
	if err != nil {
		logger.Error("failed while make http request:", err)
		return "", "", errors.NewInternalServerErr()
	}

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		logger.Error("failed to decode user info:", err)
		return "", "", errors.NewInternalServerErr()
	}

	id := userInfo["id"].(string)
	name := userInfo["name"].(string)
	email := userInfo["email"].(string)

	account := uc.repo.FindByEmail(email)

	if account == nil {
		account = entities.NewAccount(name, email, "", "", 23, id)

		uc.repo.Create(account)
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

	country, city, err := uc.glProvider.GetInfo(data.IP)
	if err != nil {
		logger.Error("Error trying to retrive geolocation data", err)
		return "", "", errors.NewInternalServerErr()
	}

	if lastCountry != "" && (lastCountry != country || lastCity != city) {
		msg := "Sua conta foi acessada em outra localização. caso não tenha sido você recomendamos que altere sua senha. obrigado pela atenção!"
		uc.emailProvider.SendMail(account.GetEmail(), "login suspeito.", msg)
	}

	deviceDetails := utils.GetDeviceDetails(data.UserAgent)
	now := time.Now()

	device := uc.devicesRepo.FindByIpAndAccountId(data.IP, account.GetID())

	if device == nil {
		device = entities.NewDevice(
			account.GetID(),
			deviceDetails.Name,
			data.UserAgent,
			deviceDetails.Platform,
			data.IP,
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
	account.SetLastLoginIp(data.IP)

	err = uc.repo.Update(account)
	if err != nil {
		logger.Error("Error trying to update account data.", err)
		return "", "", errors.NewInternalServerErr()
	}

	return accessToken, refreshToken, nil
}
