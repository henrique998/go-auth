package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVerificationCode_NewVerificationCode(t *testing.T) {
	assert := assert.New(t)

	value := "sample_code_value"
	accountId := "account123"
	expiresAt := time.Now().Add(24 * time.Hour)

	verificationCode := NewVerificationCode(value, accountId, expiresAt)

	assert.NotNil(verificationCode)
	assert.Equal(value, verificationCode.GetValue())
	assert.Equal(accountId, verificationCode.GetAccountID())
	assert.WithinDuration(expiresAt, verificationCode.GetExpiresAt(), time.Second)
	assert.WithinDuration(time.Now(), verificationCode.GetCreatedAt(), time.Second)
}
