package domain

import (
	"github.com/google/uuid"
)

type UserID string
type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleTeacher UserRole = "teacher"
	RoleStudent UserRole = "student"
)

type User struct {
	id          UserID
	email       Email
	phoneNumber PhoneNumber
	status      UserStatus
	roles       []UserRole
}

func NewUser(email Email, phone PhoneNumber) *User {
	return &User{
		id:          UserID(uuid.NewString()),
		email:       email,
		phoneNumber: phone,
		status:      StatusPending,
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

func (u *User) IsAdmin() bool {
	for _, role := range u.roles {
		if role == RoleAdmin {
			return true
		}
	}
	return false
}

func (u *User) Activate() error {
	if u == nil {
		return ErrorUserNotFound
	}

	if u.status == StatusActive {
		return ErrorUserAlreadyActive
	}
	u.status = StatusActive

	return nil
}

func (u *User) Deactivate() error {
	if u.status == StatusInactive {
		return ErrorUserAlreadyInactive
	}
	u.status = StatusInactive

	return nil
}

func (u *User) ID() UserID               { return u.id }
func (u *User) Email() Email             { return u.email }
func (u *User) PhoneNumber() PhoneNumber { return u.phoneNumber }
func (u *User) Status() UserStatus       { return u.status }
func (u *User) IsActive() bool           { return u.status == StatusActive }
