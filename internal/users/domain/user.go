package domain

import (
	"errors"

	"github.com/google/uuid"
)

type UserID string
type UserStatus string

const (
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
	StatusPending  UserStatus = "pending"
)

type User struct {
	id          UserID
	email       Email
	phoneNumber PhoneNumber
	status      UserStatus
}

func NewUser(email Email, phone PhoneNumber) *User {
	return &User{
		id:          UserID(uuid.NewString()),
		email:       email,
		phoneNumber: phone,
		status:      StatusActive,
	}
}

func RehydrateUser(id UserID, email Email, phone PhoneNumber, status UserStatus) *User {
	return &User{
		id:          id,
		email:       email,
		phoneNumber: phone,
		status:      status,
	}
}

func (u *User) Deactivate() error {
	if u.status == StatusInactive {
		return errors.New("already inactive")
	}
	u.status = StatusInactive

	return nil
}

func (u *User) ID() UserID               { return u.id }
func (u *User) Email() Email             { return u.email }
func (u *User) PhoneNumber() PhoneNumber { return u.phoneNumber }
func (u *User) Status() UserStatus       { return u.status }
func (u *User) IsActive() bool           { return u.status == StatusActive }
