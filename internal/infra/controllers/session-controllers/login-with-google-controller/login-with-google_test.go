package loginwithgooglecontroller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
	usecasesmocks "github.com/henrique998/go-auth/test/usecases-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLoginWithGoogleController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecasesmocks.NewMockLoginWithGoogleUseCase(ctrl)
	sut := NewLoginWithGoogleController(usecase)

	t.Run("It should be able to login with google", func(t *testing.T) {
		app := fiber.New()
		app.Post("/login/google", sut.Handle)

		ip := "192.168.0.1"
		userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

		data := request.LoginWithGoogleRequest{
			Code: "google-code",
		}

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("POST", "/login/google", bytes.NewBuffer(jsonData))
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
