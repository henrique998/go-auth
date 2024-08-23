package refreshtokencontroller

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

func TestRefreshTokenController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockRefreshTokenUseCase(ctrl)
	sut := NewRefreshTokenController(usecase)

	t.Run("It should be able to refresh token", func(t *testing.T) {
		app := fiber.New()
		app.Post("/login/refresh-token", sut.Handle)

		data := map[string]string{
			"refresh_token": "refresh-token-value",
		}

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("POST", "/login/refresh-token", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		usecase.EXPECT().Execute(gomock.Any()).Return("access-token", "refresh-token", nil)

		res, err := app.Test(req)
		assert.NoError(err)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Contains(string(body), "success")
	})
}
