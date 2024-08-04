package entities

import "time"

type account struct {
	id        string
	name      string
	email     string
	pass      *string
	phone     *string
	age       int8
	createdAt time.Time
	updatedAt *time.Time
}

type Account interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPass() *string
	GetPhone() *string
	GetAge() int8
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

func (u *account) GetPhone() *string {
	return u.phone
}

func (u *account) GetAge() int8 {
	return u.age
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

func NewAccount(name, email, pass, phone string, age int8) Account {
	var passVal, phoneVal *string

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

	return &account{
		id:        generateUUID(),
		name:      name,
		email:     email,
		pass:      passVal,
		phone:     phoneVal,
		age:       age,
		createdAt: time.Now(),
		updatedAt: nil,
	}
}

func NewExistingAccount(id, name, email, pass, phone string, age int8, createdAt, updatedAt time.Time) Account {
	return &account{
		id:        id,
		name:      name,
		email:     email,
		pass:      &pass,
		phone:     &phone,
		age:       age,
		createdAt: time.Now(),
		updatedAt: &updatedAt,
	}
}
