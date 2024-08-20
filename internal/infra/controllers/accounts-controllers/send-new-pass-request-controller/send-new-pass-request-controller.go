package sendnewpassrequestcontroller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type Body struct {
	Email string `json:"email"`
}

func (ctrl *sendNewPassRequestController) Handle(c fiber.Ctx) error {
	body := c.Body()

	var req Body

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err := ctrl.usecase.Execute(req.Email)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.JSON(fiber.Map{
		"message": "password request sent successfuly!",
	})
}
