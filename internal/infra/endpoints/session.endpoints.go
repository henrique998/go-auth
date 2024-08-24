package endpoints

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v3"
	loginwithcredentialsusecase "github.com/henrique998/go-auth/internal/app/usecases/session-usecases/login-with-credentials-usecase"
	loginwithgoogleusecase "github.com/henrique998/go-auth/internal/app/usecases/session-usecases/login-with-google-usecase"
	loginwithmagiclinkusecase "github.com/henrique998/go-auth/internal/app/usecases/session-usecases/login-with-magic-link-usecase"
	refreshtokenusecase "github.com/henrique998/go-auth/internal/app/usecases/session-usecases/refresh-token-usecase"
	requestmagiclinkusecase "github.com/henrique998/go-auth/internal/app/usecases/session-usecases/request-magic-link-usecase"
	loginwithcredentialscontroller "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers/login-with-credentials-controller"
	loginwithgooglecontroller "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers/login-with-google-controller"
	loginwithmagiclinkcontroller "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers/login-with-magic-link-controller"
	logoutcontroller "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers/logout-controller"
	redirectgooglelogin "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers/redirect-google-login"
	refreshtokencontroller "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers/refresh-token-controller"
	requestmagiclinkcontroller "github.com/henrique998/go-auth/internal/infra/controllers/session-controllers/request-magic-link-controller"
	"github.com/henrique998/go-auth/internal/infra/database/repositories"
	"github.com/henrique998/go-auth/internal/infra/providers"
)

func SetupSessionEndpoints(app *fiber.App, db *sql.DB) {
	accountsRepo := repositories.PGAccountsRepository{Db: db}
	devicesRepo := repositories.PGDevicesRepository{Db: db}
	laRepo := repositories.PGLoginAttemptsRepository{Db: db}
	emailProvider := providers.NewResendEmailProvider()
	rtRepo := repositories.PGRefreshTokensRepository{Db: db}
	atProvider := providers.NewAuthTokensProvider(&rtRepo)
	glProvider := providers.NewIpStackGeoLocationProvider(os.Getenv("IPSTACK_API_KEY"))
	mlRepo := repositories.PGMagicLinksRepository{Db: db}

	loginWithCredentialsUseCase := loginwithcredentialsusecase.NewLoginWithCredentialsUseCase(
		&accountsRepo,
		&devicesRepo,
		&laRepo,
		emailProvider,
		atProvider,
		glProvider,
	)
	loginWithGoogleUseCase := loginwithgoogleusecase.NewLoginWithGoogleUseCase(
		&accountsRepo,
		emailProvider,
		atProvider,
		&devicesRepo,
		glProvider,
	)
	loginWithMagicLinkUseCase := loginwithmagiclinkusecase.NewLoginWithMagicLinkUseCase(
		&accountsRepo,
		&mlRepo,
		&devicesRepo,
		atProvider,
		emailProvider,
		glProvider,
	)
	refreshTokenUseCase := refreshtokenusecase.NewRefreshTokenUseCase(
		&rtRepo,
		atProvider,
	)
	requestMagicLinkUseCase := requestmagiclinkusecase.NewRequestMagicLinkUseCase(
		&accountsRepo,
		&mlRepo,
		emailProvider,
	)

	loginWithCredentialsController := loginwithcredentialscontroller.NewLoginWithCredentialsController(loginWithCredentialsUseCase)
	loginWithGoogleController := loginwithgooglecontroller.NewLoginWithGoogleController(loginWithGoogleUseCase)
	loginWithMagicLinkController := loginwithmagiclinkcontroller.NewLoginWithMagicLinkController(loginWithMagicLinkUseCase)
	refreshTokenController := refreshtokencontroller.NewRefreshTokenController(refreshTokenUseCase)
	requestMagicLinkController := requestmagiclinkcontroller.NewRequestMagicLinkController(requestMagicLinkUseCase)

	app.Post("/login/credentials", loginWithCredentialsController.Handle)
	app.Post("/login/google", loginWithGoogleController.Handle)
	app.Post("/login/magic-link", loginWithMagicLinkController.Handle)
	app.Post("/logout", logoutcontroller.LogoutController)
	app.Get("/login/google/redirect", redirectgooglelogin.RedirectGoogleLogin)
	app.Get("/login/callback/google", func(c fiber.Ctx) error {
		code := c.Query("code")

		codeMap := map[string]string{
			"code": code,
		}

		return c.JSON(codeMap)
	})
	app.Post("/login/refresh-token", refreshTokenController.Handle)
	app.Post("/login/magic-link/request", requestMagicLinkController.Handle)
}
