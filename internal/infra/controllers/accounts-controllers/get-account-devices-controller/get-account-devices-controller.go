package getaccountdevicescontroller

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (gc *getAccountDevicesController) Handle(c fiber.Ctx) error {
	cookie := c.Cookies("goauth:access_token")

	accountId, err := utils.ParseJWTToken(cookie, os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	devices, err := gc.usecase.Execute(accountId)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(err.GetMessage())
	}

	return c.JSON(fiber.Map{
		"devices": devices,
	})
}
