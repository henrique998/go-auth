package sendnewpassrequestcontroller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	usecasesmocks "github.com/henrique998/go-auth/test/usecases-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSendNewPassRequestController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockSendNewPassRequestUseCase(ctrl)
	sut := NewSendNewPassRequestController(usecase)

	t.Run("It should be able to send new pass request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/accounts/send-new-pass-request", sut.Handle)

		data := map[string]string{
			"email": "jhondoe@gmail.com",
		}

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("POST", "/accounts/send-new-pass-request", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		usecase.EXPECT().Execute(data["email"]).Return(nil)

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Contains(string(body), "password request sent successfuly!")
	})
}
