package verify2facodecontroller

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func (ctrl *verify2FACodeController) Handle(c fiber.Ctx) error {
	cookie := c.Cookies("goauth:access_token")

	accountId, err := utils.ParseJWTToken(cookie, os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(err.GetStatus()).JSON(fiber.Map{
			"error": err.GetMessage(),
		})
	}

	body := c.Body()

	var req request.Verify2faCodeRequest

	req.AccountId = accountId

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = ctrl.usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(fiber.Map{
			"error": err.GetMessage(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Two factor authentication done succesfully!",
	})
}
