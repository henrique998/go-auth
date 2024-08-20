package endpoints

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	createaccountusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/create-account-usecase"
	getaccountdevicesusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/get-account-devices-usecase"
	createaccountcontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/create-account-controller"
	getaccountdevicescontroller "github.com/henrique998/go-auth/internal/infra/controllers/accounts-controllers/get-account-devices-controller"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
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

	createAccountUsecase := createaccountusecase.NewCreateAccountUseCase(
		&accountsRepo,
		&vcRepo,
		emailProvider,
	)
	getAccountDevicesUseCase := getaccountdevicesusecase.NewGetAccountDevicesUseCase(
		&accountsRepo,
		&devicesRepo,
	)
	createAccountController := createaccountcontroller.NewCreateAccountController(createAccountUsecase)
	getAccountDevicesController := getaccountdevicescontroller.NewGetAccountDevicesController(getAccountDevicesUseCase)

	app.Post("/accounts", createAccountController.Handle)
	app.Get("/accounts/devices", getAccountDevicesController.Handle)
}
