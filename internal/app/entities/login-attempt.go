package entities

import "time"

type loginAttempt struct {
	id          string
	email       string
	ip          string
	userAgent   string
	success     bool
	attemptedAt time.Time
}

type LoginAttempt interface {
	GetID() string
	GetEmail() string
	GetIP() string
	GetUserAgent() string
	GetSuccess() bool
	GetAttemptedAt() time.Time
}

func (l *loginAttempt) GetID() string {
	return l.id
}

func (l *loginAttempt) GetEmail() string {
	return l.email
}

func (l *loginAttempt) GetIP() string {
	return l.ip
}

func (l *loginAttempt) GetUserAgent() string {
	return l.userAgent
}

func (l *loginAttempt) GetSuccess() bool {
	return l.success
}

func (l *loginAttempt) GetAttemptedAt() time.Time {
	return l.attemptedAt
}

func NewLoginAttempt(email, ip, userAgent string, success bool) LoginAttempt {
	return &loginAttempt{
		id:          generateUUID(),
		email:       email,
		ip:          ip,
		userAgent:   userAgent,
		success:     success,
		attemptedAt: time.Now(),
	}
}

func NewExistingLoginAttempt(id, email, ip, userAgent string, success bool, attemptedAt time.Time) LoginAttempt {
	return &loginAttempt{
		id:          id,
		email:       email,
		ip:          ip,
		userAgent:   userAgent,
		success:     success,
		attemptedAt: attemptedAt,
	}
}
