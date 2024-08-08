package createaccountusecase

import (
	"fmt"
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (uc *createAccountUsecase) Execute(req request.CreateAccountRequest) appError.AppErr {
	logger.Info("Init CreateAccount UseCase")

	account := uc.repo.FindByEmail(req.Email)

	if account != nil {
		return appError.NewBadRequestErr("account already exists")
	}

	pass_hash, passErr := utils.HashPass(req.Pass)
	if passErr != nil {
		return appError.NewInternalServerErr()
	}

	data := entities.NewAccount(req.Name, req.Email, pass_hash, req.Phone, req.Age, "")

	err := uc.repo.Create(data)
	if err != nil {
		logger.Error("Error trying to create account.", err)
		return appError.NewInternalServerErr()
	}

	codeString, codeErr := utils.GenerateCode(10)
	if codeErr != nil {
		logger.Error("Error trying to generate code.", passErr)
		return appError.NewInternalServerErr()
	}

	expiresAt := time.Now().Add(time.Hour * 2)

	verificationCode := entities.NewVerificationCode(codeString, data.GetID(), expiresAt)

	err = uc.vcRepo.Create(verificationCode)
	if err != nil {
		logger.Error("Error trying to create verification code.", err)
		return appError.NewInternalServerErr()
	}

	appBaseUrl := os.Getenv("BASE_URL")
	verificationUrl := fmt.Sprintf("%saccounts/verify-email?code=%s", appBaseUrl, codeString)

	body := fmt.Sprintf(`Olá, 
	
	Por favor, verifique seu endereço de e-mail clicando no link abaixo:
	
	%s
	
	Se você não se cadastrou em nosso site, ignore este e-mail.
	
	Obrigado!`, verificationUrl)

	emailErr := uc.emailProvider.SendMail(req.Email, "Account verification.", body)
	if emailErr != nil {
		logger.Error("Error trying to send verification email.", emailErr)
		return appError.NewInternalServerErr()
	}

	return nil
}
