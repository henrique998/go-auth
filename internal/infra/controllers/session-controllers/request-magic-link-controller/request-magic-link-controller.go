package requestmagiclinkcontroller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type Request struct {
	Email string `json:"email"`
}

func (ctrl *requestMagicLinkController) Handle(c fiber.Ctx) error {
	body := c.Body()

	var req Request

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	err := ctrl.usecase.Execute(req.Email)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(fiber.Map{
			"error": err.GetMessage(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "magic link sent successfuly!",
	})
}
