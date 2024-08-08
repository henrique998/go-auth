package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMagicLink_NewMagicLink(t *testing.T) {
	assert := assert.New(t)

	accountId := "account123"
	code := "ABC123"
	expiresAt := time.Now().Add(24 * time.Hour)

	magicLink := NewMagicLink(accountId, code, expiresAt)

	assert.NotNil(magicLink)
	assert.Equal(accountId, magicLink.GetAccountID())
	assert.Equal(code, magicLink.GetCode())
	assert.WithinDuration(expiresAt, magicLink.GetExpiresAt(), time.Second)
	assert.WithinDuration(time.Now(), magicLink.GetCreatedAt(), time.Second)
}
