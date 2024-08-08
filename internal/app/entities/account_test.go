package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccount_NewAccount(t *testing.T) {
	assert := assert.New(t)

	name := "jhon"
	email := "jhondoe@gmail.com"
	pass := "123456"
	phone := "999999999"
	var age int8 = 23

	account := NewAccount(name, email, pass, phone, age, "")

	assert.NotNil(account)
	assert.Equal(name, account.GetName())
	assert.Equal(email, account.GetEmail())
	assert.Equal(pass, *account.GetPass())
	assert.Equal(phone, *account.GetPhone())
	assert.Equal(age, account.GetAge())
}

func TestAccount_NewExistingAccount(t *testing.T) {
	assert := assert.New(t)

	id := "existing-id"
	name := "jhon"
	email := "jhondoe@gmail.com"
	pass := "123456"
	phone := "999999999"
	var age int8 = 23
	now := time.Now()
	lastLoginAt := now.Add(-1 * time.Hour)
	lastLoginIp := "0.0.0.0"
	lastLoginCountry := "br"
	lastLoginCity := "sp"

	account := NewExistingAccount(
		id,
		name,
		email,
		pass,
		phone,
		age,
		"",
		true,
		false,
		lastLoginAt,
		lastLoginIp,
		lastLoginCountry,
		lastLoginCity,
		now,
		now,
	)

	assert.NotNil(account)
	assert.Equal(id, account.GetID())
	assert.Equal(name, account.GetName())
	assert.Equal(email, account.GetEmail())
	assert.Equal(pass, *account.GetPass())
	assert.Equal(phone, *account.GetPhone())
	assert.True(account.GetIs2FaEnabled())
	assert.False(account.GetIsEmailVerified())
}
