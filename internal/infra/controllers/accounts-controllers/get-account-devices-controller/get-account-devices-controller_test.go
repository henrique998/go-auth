package getaccountdevicescontroller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/infra/endpoints/middlewares"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/henrique998/go-auth/test/mocks"
	usecasesmocks "github.com/henrique998/go-auth/test/usecases-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAccountDevicesController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockGetAccountDevicesUsecase(ctrl)
	sut := NewGetAccountDevicesController(usecase)

	t.Run("It should return devices when valid token is provided", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "", 23, "")

		mockRepo := mocks.NewMockAccountsRepository(ctrl)
		mockRepo.EXPECT().FindById(account.GetID()).Return(account)

		app := fiber.New()
		app.Use(middlewares.AuthMiddleware(mockRepo))
		app.Get("/accounts/devices", func(c fiber.Ctx) error {
			return sut.Handle(c)
		})

		req := httptest.NewRequest("GET", "/accounts/devices", nil)
		expiresAt := time.Now().Add(15 * time.Minute)
		tokenStr, _ := utils.GenerateJWTToken(account.GetID(), expiresAt, os.Getenv("JWT_SECRET"))
		cookieValue := fmt.Sprintf("goauth:access_token=%s; Expires=%s; HttpOnly; Path=/",
			tokenStr, expiresAt.Format(time.RFC1123))
		req.Header.Set("Cookie", cookieValue)

		devices := []entities.Device{
			entities.NewDevice("account-id", "iphone", "Mozilla/5.0", "iOS", "192.168.0.1", time.Now().Add(-1*time.Hour)),
		}

		usecase.EXPECT().Execute(account.GetID()).Return(devices, nil)

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)
		var response map[string][]entities.Device
		json.Unmarshal(body, &response)

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Len(response["devices"], 1)
	})
}
