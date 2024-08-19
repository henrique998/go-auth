package middlewares

import (
	"errors"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

func AuthMiddleware(repo contracts.AccountsRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		accessTokenStr := c.Cookies("goauth:access_token")

		if accessTokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - Token not present",
			})
		}

		accountId, err := utils.ParseJWTToken(accessTokenStr, os.Getenv("JWT_SECRET"))
		if err != nil {
			logger.Error("Parse token error", errors.New(err.GetMessage()))
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token is malformed",
			})
		}

		account := repo.FindById(accountId)

		if account == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - Account not found",
			})
		}

		return c.Next()
	}
}
