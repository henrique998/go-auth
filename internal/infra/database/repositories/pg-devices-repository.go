package repositories

import (
	"database/sql"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGDevicesRepository struct {
	Db *sql.DB
}

type deviceRecord struct {
	id          string
	accountID   string
	deviceName  string
	userAgent   string
	platform    string
	ip          string
	createdAt   time.Time
	updatedAt   time.Time
	lastLoginAt time.Time
}

func (r *PGDevicesRepository) FindByIpAndAccountId(ip, accountId string) entities.Device {
	var deviceData deviceRecord

	query := "SELECT * FROM devices WHERE ip = $1 AND account_id = $2 LIMIT 1"

	row := r.Db.QueryRow(query, ip, accountId)

	err := row.Scan(
		&deviceData.id,
		&deviceData.accountID,
		&deviceData.deviceName,
		&deviceData.userAgent,
		&deviceData.platform,
		&deviceData.ip,
		&deviceData.createdAt,
		&deviceData.updatedAt,
		&deviceData.lastLoginAt,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error("Error trying to retrive device data", err)
		}
		return nil
	}

	device := entities.NewExistingDevice(
		deviceData.id,
		deviceData.accountID,
		deviceData.deviceName,
		deviceData.userAgent,
		deviceData.platform,
		deviceData.ip,
		deviceData.createdAt,
		deviceData.lastLoginAt,
		deviceData.updatedAt,
	)

	return device
}

func (r *PGDevicesRepository) FindManyByAccountId(accountId string) []entities.Device {
	query := "SELECT id, account_id, device_name, user_agent, platform, ip, created_at, updated_at, last_login_at FROM devices WHERE account_id = $1"

	rows, err := r.Db.Query(query, accountId)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var devices []entities.Device

	for rows.Next() {
		var deviceData deviceRecord

		err := rows.Scan(
			&deviceData.id,
			&deviceData.accountID,
			&deviceData.deviceName,
			&deviceData.userAgent,
			&deviceData.platform,
			&deviceData.ip,
			&deviceData.createdAt,
			&deviceData.updatedAt,
			&deviceData.lastLoginAt,
		)
		if err != nil {
			return nil
		}

		device := entities.NewExistingDevice(
			deviceData.id,
			deviceData.accountID,
			deviceData.deviceName,
			deviceData.userAgent,
			deviceData.platform,
			deviceData.ip,
			deviceData.createdAt,
			deviceData.lastLoginAt,
			deviceData.updatedAt,
		)

		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return devices
}

func (r *PGDevicesRepository) Create(device entities.Device) error {
	query :=
		`INSERT INTO devices (id, account_id, device_name, user_agent, platform, ip, created_at, updated_at, last_login_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.Db.Exec(query,
		device.GetID(),
		device.GetAccountID(),
		device.GetDeviceName(),
		device.GetUserAgent(),
		device.GetPlatform(),
		device.GetIP(),
		device.GetCreatedAt(),
		device.GetUpdatedAt(),
		device.GetLastLoginAt(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PGDevicesRepository) Update(device entities.Device) error {
	query := "UPDATE devices SET device_name = $1, updated_at = $2 WHERE id = $3"

	_, err := r.Db.Exec(
		query,
		device.GetDeviceName(),
		device.GetUpdatedAt(),
		device.GetID(),
	)
	if err != nil {
		return err
	}

	return nil
}
