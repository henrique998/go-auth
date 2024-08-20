package updatepasscontroller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
)

func (ctrl *updatePassController) Handle(c fiber.Ctx) error {
	body := c.Body()

	var req request.NewPassRequest

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err := ctrl.usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.JSON(fiber.Map{
		"message": "password updated successfuly!",
	})
}
