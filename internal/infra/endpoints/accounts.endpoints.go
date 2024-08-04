package endpoints

import (
	"github.com/gofiber/fiber/v3"
	createaccountusecase "github.com/henrique998/go-auth-2/internal/app/usecases/create-account-usecase"
	createaccountcontroller "github.com/henrique998/go-auth-2/internal/infra/controllers/create-account-controller"
	"github.com/henrique998/go-auth-2/internal/infra/database/repositories"
)

func accountsEndpoints(app *fiber.App) {
	accountsRepo := repositories.PGAccountsRepository{
		Db: nil,
	}

	createaccountusecase := createaccountusecase.NewCreateAccountUseCase(&accountsRepo)
	createAccountController := createaccountcontroller.NewCreateAccountController(createaccountusecase)

	app.Get("/accounts", createAccountController.Handle)
}
