package entities

import "time"

type magicLink struct {
	id        string
	accountId string
	code      string
	expiresAt time.Time
	createdAt time.Time
}

type MagicLink interface {
	GetID() string
	GetAccountID() string
	GetCode() string
	GetExpiresAt() time.Time
	GetCreatedAt() time.Time
}

func (m *magicLink) GetID() string {
	return m.id
}

func (m *magicLink) GetAccountID() string {
	return m.accountId
}

func (m *magicLink) GetCode() string {
	return m.code
}

func (m *magicLink) GetExpiresAt() time.Time {
	return m.expiresAt
}

func (m *magicLink) GetCreatedAt() time.Time {
	return m.createdAt
}

func NewMagicLink(accountId, code string, expiresAt time.Time) MagicLink {
	return &magicLink{
		id:        generateUUID(),
		accountId: accountId,
		code:      code,
		expiresAt: expiresAt,
		createdAt: time.Now(),
	}
}

func NewExistingMagicLink(id, accountId, code string, expiresAt, createdAt time.Time) MagicLink {
	return &magicLink{
		id:        id,
		accountId: accountId,
		code:      code,
		expiresAt: expiresAt,
		createdAt: createdAt,
	}
}
