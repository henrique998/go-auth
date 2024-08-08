package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoginAttempt_NewLoginAttempt(t *testing.T) {
	assert := assert.New(t)

	email := "test@example.com"
	ip := "192.168.0.1"
	userAgent := "Mozilla/5.0"
	success := true

	loginAttempt := NewLoginAttempt(email, ip, userAgent, success)

	assert.NotNil(loginAttempt)
	assert.Equal(email, loginAttempt.GetEmail())
	assert.Equal(ip, loginAttempt.GetIP())
	assert.Equal(userAgent, loginAttempt.GetUserAgent())
	assert.Equal(success, loginAttempt.GetSuccess())
	assert.WithinDuration(time.Now(), loginAttempt.GetAttemptedAt(), time.Second)
}
