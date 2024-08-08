package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRefreshToken_NewRefreshToken(t *testing.T) {
	assert := assert.New(t)

	value := "sample_token_value"
	accountId := "account123"
	expiresAt := time.Now().Add(24 * time.Hour)

	refreshToken := NewRefreshToken(value, accountId, expiresAt)

	assert.NotNil(refreshToken)
	assert.Equal(value, refreshToken.GetValue())
	assert.Equal(accountId, refreshToken.GetAccountID())
	assert.WithinDuration(expiresAt, refreshToken.GetExpiresAt(), time.Second)
	assert.WithinDuration(time.Now(), refreshToken.GetCreatedAt(), time.Second)
}
