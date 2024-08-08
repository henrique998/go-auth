package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/stretchr/testify/assert"
)

func TestPGDevicesRepository_FindByIpAndAccountId(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	deviceId := "fake-device-id"
	accountId := "fake-account-id"
	ip := "fake-device-ip"

	deviceData := entities.NewExistingDevice(
		deviceId,
		accountId,
		"device-name",
		"user-agent",
		"IOS",
		ip,
		time.Now().Add(-2*(time.Hour*24*7)),
		time.Now().Add(-1*time.Hour),
		time.Now(),
	)

	repo := PGDevicesRepository{
		Db: db,
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "account_id", "device_name", "user_agent", "platform", "ip_address", "created_at", "updated_at", "last_login_at",
		}).AddRow(
			deviceData.GetID(),
			deviceData.GetAccountID(),
			deviceData.GetDeviceName(),
			deviceData.GetUserAgent(),
			deviceData.GetPlatform(),
			deviceData.GetIP(),
			deviceData.GetCreatedAt(),
			deviceData.GetUpdatedAt(),
			deviceData.GetLastLoginAt(),
		)

		mock.ExpectQuery(`SELECT \* FROM devices WHERE ip = \$1 AND account_id = \$2 LIMIT 1`).
			WithArgs(ip, accountId).
			WillReturnRows(rows)

		device := repo.FindByIpAndAccountId(ip, accountId)

		assert.NotNil(device)
		assert.Equal(accountId, device.GetAccountID())
		assert.Equal(ip, device.GetIP())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM devices WHERE ip = \$1 AND account_id = \$2 LIMIT 1`).
			WithArgs(ip, accountId).
			WillReturnError(errors.New("some other error"))

		device := repo.FindByIpAndAccountId(ip, accountId)

		assert.Nil(device)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGDevicesRepository_FindManyByAccountId(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	accountId := "fake-account-id"

	deviceData := entities.NewExistingDevice(
		"fake-device-id",
		accountId,
		"device-name",
		"user-agent",
		"IOS",
		"fake-ip",
		time.Now().Add(-2*(time.Hour*24*7)),
		time.Now().Add(-1*time.Hour),
		time.Now(),
	)

	repo := PGDevicesRepository{
		Db: db,
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "account_id", "device_name", "user_agent", "platform", "ip", "created_at", "updated_at", "last_login_at",
		}).AddRow(
			deviceData.GetID(),
			deviceData.GetAccountID(),
			deviceData.GetDeviceName(),
			deviceData.GetUserAgent(),
			deviceData.GetPlatform(),
			deviceData.GetIP(),
			deviceData.GetCreatedAt(),
			deviceData.GetUpdatedAt(),
			deviceData.GetLastLoginAt(),
		)

		mock.ExpectQuery(`SELECT id, account_id, device_name, user_agent, platform, ip, created_at, updated_at, last_login_at FROM devices WHERE account_id = \$1`).
			WithArgs(accountId).
			WillReturnRows(rows)

		devices := repo.FindManyByAccountId(accountId)

		assert.Len(devices, 1)
		assert.Equal(devices[0].GetID(), "fake-device-id")

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGDevicesRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGDevicesRepository{Db: db}

	deviceData := entities.NewExistingDevice(
		"fake-device-id",
		"fake-account-id",
		"device-name",
		"user-agent",
		"IOS",
		"fake-ip",
		time.Now().Add(-2*(time.Hour*24*7)),
		time.Now().Add(-1*time.Hour),
		time.Now(),
	)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO devices \(id, account_id, device_name, user_agent, platform, ip, created_at, updated_at, last_login_at\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9\)`).
			WithArgs(
				deviceData.GetID(),
				deviceData.GetAccountID(),
				deviceData.GetDeviceName(),
				deviceData.GetUserAgent(),
				deviceData.GetPlatform(),
				deviceData.GetIP(),
				deviceData.GetCreatedAt(),
				deviceData.GetUpdatedAt(),
				deviceData.GetLastLoginAt(),
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Create(deviceData)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO devices \(id, account_id, device_name, user_agent, platform, ip, created_at, updated_at, last_login_at\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9\)`).
			WithArgs(
				deviceData.GetID(),
				deviceData.GetAccountID(),
				deviceData.GetDeviceName(),
				deviceData.GetUserAgent(),
				deviceData.GetPlatform(),
				deviceData.GetIP(),
				deviceData.GetCreatedAt(),
				deviceData.GetUpdatedAt(),
				deviceData.GetLastLoginAt(),
			).WillReturnError(errors.New("insert failed"))

		err = repo.Create(deviceData)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGDevicesRepository_Update(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGDevicesRepository{Db: db}

	deviceData := entities.NewExistingDevice(
		"fake-device-id",
		"fake-account-id",
		"device-name",
		"user-agent",
		"IOS",
		"fake-ip",
		time.Now().Add(-2*(time.Hour*24*7)),
		time.Now().Add(-1*time.Hour),
		time.Now(),
	)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`UPDATE devices SET device_name = \$1, updated_at = \$2 WHERE id = \$3`).
			WithArgs(
				deviceData.GetDeviceName(),
				deviceData.GetUpdatedAt(),
				deviceData.GetID(),
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Update(deviceData)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`UPDATE devices SET device_name = \$1, updated_at = \$2 WHERE id = \$3`).
			WithArgs(
				deviceData.GetDeviceName(),
				deviceData.GetUpdatedAt(),
				deviceData.GetID(),
			).WillReturnError(errors.New("update failed"))

		err = repo.Update(deviceData)

		assert.Error(err)
		assert.Equal("update failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}
