package entities

import "time"

type account struct {
	id               string
	name             string
	email            string
	pass             *string
	phone            *string
	age              int8
	providerId       *string
	is2faEnabled     bool
	isEmailVerified  bool
	lastLoginAt      *time.Time
	lastLoginIp      *string
	lastLoginCountry *string
	lastLoginCity    *string
	createdAt        time.Time
	updatedAt        *time.Time
}

type Account interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPass() *string
	SetPass(pass string)
	GetPhone() *string
	GetAge() int8
	GetProviderId() *string
	GetIs2FaEnabled() bool
	Enable2FA()
	GetIsEmailVerified() bool
	VerifyEmail()
	GetLastLoginAt() *time.Time
	LoginNow()
	GetLastLoginIp() *string
	SetLastLoginIp(ip string)
	GetLastLoginCountry() *string
	SetLastLoginCountry(country string)
	GetLastLoginCity() *string
	SetLastLoginCity(city string)
	GetCreatedAt() time.Time
	GetUpdatedAt() *time.Time
	Touch()
}

func (u *account) GetID() string {
	return u.id
}

func (u *account) GetName() string {
	return u.name
}

func (u *account) GetEmail() string {
	return u.email
}

func (u *account) GetPass() *string {
	return u.pass
}

func (a *account) SetPass(pass string) {
	a.pass = &pass
}

func (u *account) GetPhone() *string {
	return u.phone
}

func (u *account) GetAge() int8 {
	return u.age
}

func (u *account) GetProviderId() *string {
	return u.providerId
}

func (u *account) GetIs2FaEnabled() bool {
	return u.is2faEnabled
}

func (a *account) Enable2FA() {
	a.is2faEnabled = true
}

func (u *account) GetIsEmailVerified() bool {
	return u.isEmailVerified
}

func (a *account) VerifyEmail() {
	a.isEmailVerified = true
}

func (u *account) GetLastLoginAt() *time.Time {
	return u.lastLoginAt
}

func (a *account) LoginNow() {
	now := time.Now()
	a.lastLoginAt = &now
}

func (a *account) GetLastLoginIp() *string {
	return a.lastLoginIp
}

func (a *account) SetLastLoginIp(ip string) {
	a.lastLoginIp = &ip
}

func (u *account) GetLastLoginCountry() *string {
	return u.lastLoginCountry
}

func (a *account) SetLastLoginCountry(country string) {
	a.lastLoginCountry = &country
}

func (u *account) GetLastLoginCity() *string {
	return u.lastLoginCity
}

func (a *account) SetLastLoginCity(city string) {
	a.lastLoginCity = &city
}

func (u *account) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *account) GetUpdatedAt() *time.Time {
	return u.updatedAt
}

func (u *account) Touch() {
	now := time.Now()
	u.updatedAt = &now
}

func NewAccount(name, email, pass, phone string, age int8, providerId string) Account {
	var passVal, phoneVal, providerIdVal *string

	if pass == "" {
		passVal = nil
	} else {
		passVal = &pass
	}

	if phone == "" {
		phoneVal = nil
	} else {
		phoneVal = &phone
	}

	if providerId == "" {
		providerIdVal = nil
	} else {
		providerIdVal = &providerId
	}

	return &account{
		id:               generateUUID(),
		name:             name,
		email:            email,
		pass:             passVal,
		phone:            phoneVal,
		age:              age,
		providerId:       providerIdVal,
		is2faEnabled:     false,
		isEmailVerified:  false,
		lastLoginAt:      nil,
		lastLoginIp:      nil,
		lastLoginCountry: nil,
		lastLoginCity:    nil,
		createdAt:        time.Now(),
		updatedAt:        nil,
	}
}

func NewExistingAccount(
	id,
	name,
	email,
	pass,
	phone string,
	age int8,
	providerId *string,
	is2faEnbaled,
	isEmailVerified bool,
	lastLoginAt *time.Time,
	lastLoginIp,
	lastLoginCountry,
	lastLoginCity *string,
	createdAt time.Time,
	updatedAt *time.Time,
) Account {
	return &account{
		id:               id,
		name:             name,
		email:            email,
		pass:             &pass,
		phone:            &phone,
		age:              age,
		providerId:       providerId,
		is2faEnabled:     is2faEnbaled,
		isEmailVerified:  isEmailVerified,
		lastLoginAt:      lastLoginAt,
		lastLoginIp:      lastLoginIp,
		lastLoginCountry: lastLoginCountry,
		lastLoginCity:    lastLoginCity,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
	}
}
