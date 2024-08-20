package updatepasscontroller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	usecasesmocks "github.com/henrique998/go-auth/test/usecases-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdatePassController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockUpdatePassUseCase(ctrl)
	sut := NewUpdatePassController(usecase)

	t.Run("It should be able to update pass", func(t *testing.T) {
		app := fiber.New()
		app.Patch("/accounts/update-pass", sut.Handle)

		data := request.NewPassRequest{
			Code:                "code",
			NewPass:             "new-pass",
			NewPassConfirmation: "new-pass",
		}

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("PATCH", "/accounts/update-pass", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		usecase.EXPECT().Execute(data).Return(nil)

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Contains(string(body), "password updated successfuly!")
	})
}
