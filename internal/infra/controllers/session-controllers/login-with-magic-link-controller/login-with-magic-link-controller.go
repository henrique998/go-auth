package loginwithmagiclinkcontroller

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
)

func (ctrl *loginWithMagicLinkController) Handle(c fiber.Ctx) error {
	code := c.Query("code")

	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "code is required"})
	}

	ip := c.IP()
	userAgent := c.Get("User-Agent")

	var req request.LoginWithMagicLinkRequest
	req.Code = code
	req.IP = ip
	req.UserAgent = userAgent

	accessToken, refreshToken, err := ctrl.usecase.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(fiber.Map{
			"error": err.GetMessage(),
		})
	}

	accessTokenCookie := fiber.Cookie{
		Name:     "goauth:access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Path:     "/",
	}

	refreshTokenCookie := fiber.Cookie{
		Name:     "goauth:refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		Path:     "/",
	}

	c.Cookie(&accessTokenCookie)
	c.Cookie(&refreshTokenCookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
