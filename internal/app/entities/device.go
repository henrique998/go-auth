package entities

import "time"

type device struct {
	id          string
	accountID   string
	deviceName  string
	userAgent   string
	platform    string
	ip          string
	createdAt   time.Time
	updatedAt   *time.Time
	lastLoginAt time.Time
}

type Device interface {
	GetID() string
	GetAccountID() string
	GetDeviceName() string
	GetUserAgent() string
	GetPlatform() string
	GetIP() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetLastLoginAt() time.Time
	Touch()
}

func (d *device) GetID() string {
	return d.id
}

func (d *device) GetAccountID() string {
	return d.accountID
}

func (d *device) GetDeviceName() string {
	return d.deviceName
}

func (d *device) GetUserAgent() string {
	return d.userAgent
}

func (d *device) GetPlatform() string {
	return d.platform
}

func (d *device) GetIP() string {
	return d.ip
}

func (d *device) GetLastLoginAt() time.Time {
	return d.lastLoginAt
}

func (d *device) GetCreatedAt() time.Time {
	return d.createdAt
}

func (d *device) GetUpdatedAt() time.Time {
	return *d.updatedAt
}

func (d *device) Touch() {
	now := time.Now()
	d.updatedAt = &now
}

func NewDevice(accountId, deviceName, userAgent, platform, ip string, lastLoginAt time.Time) Device {
	return &device{
		id:          generateUUID(),
		accountID:   accountId,
		deviceName:  deviceName,
		userAgent:   userAgent,
		platform:    platform,
		ip:          ip,
		lastLoginAt: lastLoginAt,
		createdAt:   time.Now(),
		updatedAt:   nil,
	}
}

func NewExistingDevice(id, accountId, deviceName, userAgent, platform, ip string, createdAt, lastLoginAt, updatedtAt time.Time) Device {
	return &device{
		id:          id,
		accountID:   accountId,
		deviceName:  deviceName,
		userAgent:   userAgent,
		platform:    platform,
		ip:          ip,
		createdAt:   createdAt,
		updatedAt:   &updatedtAt,
		lastLoginAt: lastLoginAt,
	}
}
