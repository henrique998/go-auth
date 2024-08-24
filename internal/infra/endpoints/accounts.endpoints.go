package endpoints

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v3"
	createaccountusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/create-account-usecase"
	getaccountdevicesusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/get-account-devices-usecase"
	send2facodeusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/send-2fa-code-usecase"
	sendnewpassrequestusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/send-new-pass-request-usecase"
	updatepassusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/update-pass-usecase"
	verify2facodeusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/verify-2fa-code-usecase"
	verifyemailaccountusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/verify-email-usecase"
	createaccountcontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/create-account-controller"
	getaccountdevicescontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/get-account-devices-controller"
	send2facodecontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/send-2fa-code-controller"
	sendnewpassrequestcontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/send-new-pass-request-controller"
	updatepasscontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/update-pass-controller"
	verify2facodecontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/verify-2fa-code-controller"
	verifyemailcontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/verify-email-controller"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
	"github.com/twilio/twilio-go"
)

func SetupAccountsEndpoints(app *fiber.App, db *sql.DB) {
	accountsRepo := repositories.PGAccountsRepository{
		Db: db,
	}
	vcRepo := repositories.PGVerificationCodesRepository{
		Db: db,
	}
	devicesRepo := repositories.PGDevicesRepository{
		Db: db,
	}
	emailProvider := providers.NewResendEmailProvider()
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	twoFactorAuthProvider := providers.NewTwilioTwoFactorAuthProvider(twilioClient)

	createAccountUsecase := createaccountusecase.NewCreateAccountUseCase(
		&accountsRepo,
		&vcRepo,
		emailProvider,
	)
	getAccountDevicesUseCase := getaccountdevicesusecase.NewGetAccountDevicesUseCase(
		&accountsRepo,
		&devicesRepo,
	)
	send2FACodeUseCase := send2facodeusecase.NewSend2faCodeUseCase(
		&accountsRepo,
		&vcRepo,
		twoFactorAuthProvider,
	)
	sendNewPassUseCase := sendnewpassrequestusecase.NewSendNewPassRequestUseCase(
		&accountsRepo,
		&vcRepo,
		emailProvider,
	)
	updatePassUseCase := updatepassusecase.NewUpdatePassUseCase(
		&accountsRepo,
		&vcRepo,
	)
	verify2FACodeUseCase := verify2facodeusecase.NewVerify2faCodeUseCase(
		&accountsRepo,
		&vcRepo,
	)
	verifyEmailUseCase := verifyemailaccountusecase.NewVerifyEmailUseCase(
		&accountsRepo,
		&vcRepo,
	)

	createAccountController := createaccountcontroller.NewCreateAccountController(createAccountUsecase)
	getAccountDevicesController := getaccountdevicescontroller.NewGetAccountDevicesController(getAccountDevicesUseCase)
	send2FACodeController := send2facodecontroller.NewSend2FACodeController(send2FACodeUseCase)
	sendNewPassController := sendnewpassrequestcontroller.NewSendNewPassRequestController(sendNewPassUseCase)
	updatePassController := updatepasscontroller.NewUpdatePassController(updatePassUseCase)
	verify2FACodeController := verify2facodecontroller.NewVerify2FACodeController(verify2FACodeUseCase)
	verifyEmailController := verifyemailcontroller.NewVerifyEmailController(verifyEmailUseCase)

	app.Post("/accounts", createAccountController.Handle)
	app.Get("/accounts/devices", getAccountDevicesController.Handle)
	app.Post("/accounts/send-2fa-code", send2FACodeController.Handle)
	app.Post("/accounts/send-new-pass-request", sendNewPassController.Handle)
	app.Patch("/accounts/update-pass", updatePassController.Handle)
	app.Post("/accounts/verify-2fa-code", verify2FACodeController.Handle)
	app.Post("/accounts/verify-email", verifyEmailController.Handle)
}
