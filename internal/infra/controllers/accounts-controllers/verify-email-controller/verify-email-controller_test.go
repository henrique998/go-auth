package verifyemailcontroller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	usecasesmocks "github.com/henrique998/go-auth/test/usecases-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestVerifyEmailController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockVerifyEmailUseCase(ctrl)
	sut := NewVerifyEmailController(usecase)

	t.Run("It should be able to verify email", func(t *testing.T) {
		app := fiber.New()
		app.Get("/accounts/verify-email", sut.Handle)

		token := "valid-token"

		req := httptest.NewRequest("GET", "/accounts/verify-email?token="+token, nil)
		usecase.EXPECT().Execute(token).Return(nil)

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Contains(string(body), "Email verified successfully!")
	})
}
