package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDevice_NewDevice(t *testing.T) {
	assert := assert.New(t)

	accountId := "account-id"
	deviceName := "iPhone"
	userAgent := "Mozilla/5.0"
	platform := "iOS"
	ip := "192.168.0.1"
	lastLoginAt := time.Now().Add(-1 * time.Hour)

	device := NewDevice(accountId, deviceName, userAgent, platform, ip, lastLoginAt)

	assert.NotNil(device)
	assert.Equal(accountId, device.GetAccountID())
	assert.Equal(deviceName, device.GetDeviceName())
	assert.Equal(userAgent, device.GetUserAgent())
	assert.Equal(platform, device.GetPlatform())
	assert.Equal(ip, device.GetIP())
	assert.Equal(lastLoginAt, device.GetLastLoginAt())
	assert.WithinDuration(time.Now(), device.GetCreatedAt(), time.Second)
}

func TestDevice_NewExistingDevice(t *testing.T) {
	assert := assert.New(t)

	id := "existing-id"
	accountId := "account-id"
	deviceName := "iPhone"
	userAgent := "Mozilla/5.0"
	platform := "iOS"
	ip := "192.168.0.1"
	createdAt := time.Now().Add(-24 * time.Hour)
	lastLoginAt := time.Now().Add(-1 * time.Hour)
	updatedAt := time.Now().Add(-30 * time.Minute)

	device := NewExistingDevice(id, accountId, deviceName, userAgent, platform, ip, createdAt, lastLoginAt, updatedAt)

	assert.NotNil(device)
	assert.Equal(id, device.GetID())
	assert.Equal(accountId, device.GetAccountID())
	assert.Equal(deviceName, device.GetDeviceName())
	assert.Equal(userAgent, device.GetUserAgent())
	assert.Equal(platform, device.GetPlatform())
	assert.Equal(ip, device.GetIP())
	assert.Equal(createdAt, device.GetCreatedAt())
	assert.Equal(lastLoginAt, device.GetLastLoginAt())
	assert.Equal(updatedAt, device.GetUpdatedAt())
}
