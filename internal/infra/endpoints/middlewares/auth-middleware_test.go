package middlewares

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthMiddleware(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("It should return bad request when cookie is not present", func(t *testing.T) {
		app := fiber.New()
		app.Use(AuthMiddleware(nil))
		app.Get("/hello", func(c fiber.Ctx) error {
			return nil
		})
		req := httptest.NewRequest("GET", "/hello", nil)
		res, _ := app.Test(req)

		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusUnauthorized, res.StatusCode)
		assert.Contains(string(body), "Unauthorized - Token not present")
	})

	t.Run("It should return unauthorized when JWT token is invalid", func(t *testing.T) {
		app := fiber.New()
		app.Use(AuthMiddleware(nil))
		app.Get("/hello", func(c fiber.Ctx) error {
			return nil
		})
		req := httptest.NewRequest("GET", "/hello", nil)
		req.Header.Set("Cookie", "goauth:access_token=invalid_token")

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusUnauthorized, res.StatusCode)
		assert.Contains(string(body), "Token is malformed")
	})

	t.Run("It should return unauthorized when account was not found", func(t *testing.T) {
		mockRepo := mocks.NewMockAccountsRepository(ctrl)
		mockRepo.EXPECT().FindById(gomock.Any()).Return(nil)

		app := fiber.New()
		app.Use(AuthMiddleware(mockRepo))
		app.Get("/hello", func(c fiber.Ctx) error {
			return nil
		})
		req := httptest.NewRequest("GET", "/hello", nil)
		expiresAt := time.Now().Add(15 * time.Minute)
		tokenStr, _ := utils.GenerateJWTToken("accountId", expiresAt, os.Getenv("JWT_SECRET"))
		cookieValue := fmt.Sprintf("goauth:access_token=%s; Expires=%s; HttpOnly; Path=/",
			tokenStr, expiresAt.Format(time.RFC1123))
		req.Header.Set("Cookie", cookieValue)

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusUnauthorized, res.StatusCode)
		assert.Contains(string(body), "Unauthorized - Account not found")
	})

	t.Run("It should return body content when valid content is provided", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "", 23, "")

		mockRepo := mocks.NewMockAccountsRepository(ctrl)
		mockRepo.EXPECT().FindById(account.GetID()).Return(account)

		app := fiber.New()
		app.Use(AuthMiddleware(mockRepo))
		app.Get("/hello", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "hello world",
			})
		})

		req := httptest.NewRequest("GET", "/hello", nil)
		expiresAt := time.Now().Add(15 * time.Minute)
		tokenStr, _ := utils.GenerateJWTToken(account.GetID(), expiresAt, os.Getenv("JWT_SECRET"))
		cookieValue := fmt.Sprintf("goauth:access_token=%s; Expires=%s; HttpOnly; Path=/",
			tokenStr, expiresAt.Format(time.RFC1123))
		req.Header.Set("Cookie", cookieValue)

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Contains(string(body), "hello world")
	})
}
