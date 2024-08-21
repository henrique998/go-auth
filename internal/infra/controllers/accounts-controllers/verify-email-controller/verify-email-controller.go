package verifyemailcontroller

import "github.com/gofiber/fiber/v3"

func (ctrl *verifyEmailController) Handle(c fiber.Ctx) error {
	token := c.Query("token")

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing query parameter 'param'",
		})
	}

	err := ctrl.usecase.Execute(token)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(fiber.Map{
			"error": err.GetMessage(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Email verified successfully!",
	})
}
