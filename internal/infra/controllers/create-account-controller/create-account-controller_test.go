package createaccountcontroller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/errors"
	usecasesmocks "github.com/henrique998/go-auth/test/usecases-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateAccountController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("It should not be able to create an account that already exists", func(t *testing.T) {
		usecase := usecasesmocks.NewMockCreateAccountUsecase(ctrl)

		usecase.EXPECT().Execute(gomock.Any()).Return(errors.NewBadRequestErr("account already exists"))

		sut := NewCreateAccountController(usecase)

		app := fiber.New()
		app.Post("/accounts", sut.Handle)

		bodyData := map[string]any{
			"name":     "jhon doe",
			"email":    "jhondoe@gmail.com",
			"password": "123456",
			"phone":    "999999999",
			"age":      23,
		}

		jsonData, _ := json.Marshal(bodyData)

		req := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonData))
		res, err := app.Test(req, -1)
		assert.NoError(err)

		assert.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		usecase := usecasesmocks.NewMockCreateAccountUsecase(ctrl)

		usecase.EXPECT().Execute(gomock.Any()).Return(nil)

		sut := NewCreateAccountController(usecase)

		app := fiber.New()
		app.Post("/accounts", sut.Handle)

		bodyData := map[string]any{
			"name":     "jhon doe",
			"email":    "jhondoe@gmail.com",
			"password": "123456",
			"phone":    "999999999",
			"age":      23,
		}

		jsonData, _ := json.Marshal(bodyData)

		req := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonData))
		res, err := app.Test(req, -1)
		assert.NoError(err)

		assert.Equal(http.StatusOK, res.StatusCode)
	})
}
