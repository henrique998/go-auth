package entities

import "time"

type refreshToken struct {
	id        string
	value     string
	accountId string
	expiresAt time.Time
	createdAt time.Time
}

type RefreshToken interface {
	GetID() string
	GetValue() string
	GetAccountID() string
	GetExpiresAt() time.Time
	GetCreatedAt() time.Time
}

func (r *refreshToken) GetID() string {
	return r.id
}

func (r *refreshToken) GetValue() string {
	return r.value
}

func (r *refreshToken) GetAccountID() string {
	return r.accountId
}

func (r *refreshToken) GetExpiresAt() time.Time {
	return r.expiresAt
}

func (r *refreshToken) GetCreatedAt() time.Time {
	return r.createdAt
}

func NewRefreshToken(value, accountId string, expiresAt time.Time) RefreshToken {
	return &refreshToken{
		id:        generateUUID(),
		value:     value,
		accountId: accountId,
		expiresAt: expiresAt,
		createdAt: time.Now(),
	}
}

func NewExistingRefreshToken(id, value, accountId string, expiresAt, createdAt time.Time) RefreshToken {
	return &refreshToken{
		id:        id,
		value:     value,
		accountId: accountId,
		expiresAt: expiresAt,
		createdAt: createdAt,
	}
}
