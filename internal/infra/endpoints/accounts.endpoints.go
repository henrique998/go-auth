package endpoints

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	createaccountusecase "github.com/henrique998/go-auth/internal/app/usecases/account-usecases/create-account-usecase"
	createaccountcontroller "github.com/henrique998/go-auth/internal/infra/controllers/create-account-controller"
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
	emailProvider := providers.NewResendEmailProvider()

	createaccountusecase := createaccountusecase.NewCreateAccountUseCase(
		&accountsRepo,
		&vcRepo,
		emailProvider,
	)
	createAccountController := createaccountcontroller.NewCreateAccountController(createaccountusecase)

	app.Post("/accounts", createAccountController.Handle)
}
