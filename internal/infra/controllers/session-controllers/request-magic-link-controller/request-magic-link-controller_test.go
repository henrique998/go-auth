package requestmagiclinkcontroller

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

func TestRequestMagicLinkController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockRequestMagicLinkUseCase(ctrl)
	sut := NewRequestMagicLinkController(usecase)

	t.Run("Irt should be able to request a magic link", func(t *testing.T) {
		app := fiber.New()
		app.Post("/login/magic-link", sut.Handle)

		var data Request

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("POST", "/login/magic-link", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		usecase.EXPECT().Execute(gomock.Any()).Return(nil)

		res, err := app.Test(req)
		assert.NoError(err)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Contains(string(body), "magic link sent successfuly!")
	})
}
