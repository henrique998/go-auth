package loginwithmagiclinkcontroller

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	usecasesmocks "github.com/henrique998/go-auth/test/usecases-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLoginWithMagicLinkController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockLoginWithMagicLinkUseCase(ctrl)
	sut := NewLoginWithMagicLinkController(usecase)

	t.Run("It should return bad request if code is missing", func(t *testing.T) {
		app := fiber.New()
		app.Get("/login/magic-link", sut.Handle)

		req := httptest.NewRequest("GET", "/login/magic-link", nil)

		res, err := app.Test(req)
		assert.NoError(err)

		body, _ := io.ReadAll(res.Body)
		assert.Equal(http.StatusBadRequest, res.StatusCode)
		assert.Contains(string(body), "code is required")
	})

	t.Run("It should be able to login with magic link", func(t *testing.T) {
		app := fiber.New()
		app.Get("/login/magic-link", sut.Handle)

		code := "magic-link-code"
		ip := "192.168.0.1"
		userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

		req := httptest.NewRequest("GET", fmt.Sprintf("/login/magic-link?code=%s", code), nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", userAgent)
		req.RemoteAddr = ip + ":1234"

		usecase.EXPECT().Execute(gomock.Any()).Return("access-token", "refresh-token", nil)

		res, err := app.Test(req)
		assert.NoError(err)
		body, _ := io.ReadAll(res.Body)

		setCookieHeader := res.Header["Set-Cookie"]

		var cookieFound bool
		for _, cookie := range setCookieHeader {
			if strings.Contains(cookie, "goauth:access_token=access-token") {
				cookieFound = true
				break
			}
		}

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Contains(string(body), "success")
		assert.True(cookieFound)
	})
}
