package entities

import "time"

type verificationCode struct {
	id        string
	value     string
	accountId string
	expiresAt time.Time
	createdAt time.Time
}

type VerificationCode interface {
	GetID() string
	GetValue() string
	GetAccountID() string
	GetExpiresAt() time.Time
	GetCreatedAt() time.Time
}

func (v *verificationCode) GetID() string {
	return v.id
}

func (v *verificationCode) GetValue() string {
	return v.value
}

func (v *verificationCode) GetAccountID() string {
	return v.accountId
}

func (v *verificationCode) GetExpiresAt() time.Time {
	return v.expiresAt
}

func (v *verificationCode) GetCreatedAt() time.Time {
	return v.createdAt
}

func NewVerificationCode(value, accountId string, expiresAt time.Time) VerificationCode {
	return &verificationCode{
		id:        generateUUID(),
		value:     value,
		accountId: accountId,
		expiresAt: expiresAt,
		createdAt: time.Now(),
	}
}

func NewExistingVerificationCode(id, value, accountId string, expiresAt, createdAt time.Time) VerificationCode {
	return &verificationCode{
		id:        id,
		value:     value,
		accountId: accountId,
		expiresAt: expiresAt,
		createdAt: createdAt,
	}
}
