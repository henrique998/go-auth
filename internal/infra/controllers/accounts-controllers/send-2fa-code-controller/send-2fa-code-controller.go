package send2facodecontroller

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (sc *send2FACodeController) Handle(c fiber.Ctx) error {
	cookie := c.Cookies("goauth:access_token")

	accountId, err := utils.ParseJWTToken(cookie, os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	err = sc.usecase.Execute(accountId)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.JSON(fiber.Map{
		"message": "code sent succesfully!",
	})
}
