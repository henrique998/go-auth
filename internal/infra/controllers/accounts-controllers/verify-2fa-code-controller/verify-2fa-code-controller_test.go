package verify2facodecontroller

import (
	"bytes"
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
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/infra/endpoints/middlewares"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/henrique998/go-auth/test/mocks"
	usecasesmocks "github.com/henrique998/go-auth/test/usecases-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test2FACodeController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockVerify2faCodeUseCase(ctrl)
	sut := NewVerify2FACodeController(usecase)

	t.Run("It should be able verify 2fa code", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "", 23, "")

		mockRepo := mocks.NewMockAccountsRepository(ctrl)
		mockRepo.EXPECT().FindById(account.GetID()).Return(account)

		app := fiber.New()
		app.Use(middlewares.AuthMiddleware(mockRepo))
		app.Post("/accounts/verify-2fa-code", sut.Handle)

		data := request.Verify2faCodeRequest{
			Code:      "code",
			AccountId: account.GetID(),
		}

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("POST", "/accounts/verify-2fa-code", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		expiresAt := time.Now().Add(15 * time.Minute)
		tokenStr, _ := utils.GenerateJWTToken(account.GetID(), expiresAt, os.Getenv("JWT_SECRET"))
		cookieValue := fmt.Sprintf("goauth:access_token=%s; Expires=%s; HttpOnly; Path=/",
			tokenStr, expiresAt.Format(time.RFC1123))
		req.Header.Set("Cookie", cookieValue)

		usecase.EXPECT().Execute(data).Return(nil)

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Contains(string(body), "Two factor authentication done succesfully!")
	})
}
