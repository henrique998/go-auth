package getaccountdevicesusecase

import (
	"net/http"
	"testing"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAccountDevicesUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockDevicesRepo := mocks.NewMockDevicesRepository(ctrl)

	usecase := NewGetAccountDevicesUseCase(
		mockAccountsRepo,
		mockDevicesRepo,
	)

	t.Run("It should not be able to get account devices that not exists", func(t *testing.T) {
		accountId := "invalid-account-id"

		mockAccountsRepo.EXPECT().FindById(accountId).Return(nil)

		devices, err := usecase.Execute(accountId)

		assert.Nil(devices)
		assert.NotNil(err)
		assert.Equal("account does not exists", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should be able to get account devices", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", 23, "")

		deviceName := "device-name"
		userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
		platform := "IOS"
		ip := "0.0.0.0"

		devicesData := []entities.Device{
			entities.NewDevice(account.GetID(), deviceName, userAgent, platform, ip, time.Now()),
		}

		mockAccountsRepo.EXPECT().FindById(account.GetID()).Return(account)
		mockDevicesRepo.EXPECT().FindManyByAccountId(account.GetID()).Return(devicesData)

		devices, err := usecase.Execute(account.GetID())

		assert.Nil(err)
		assert.Len(devices, 1)
		assert.Equal(deviceName, devices[0].GetDeviceName())
	})
}
