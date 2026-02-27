package domain

import "context"

type Repository interface {
	Save(ctx context.Context, user *User) error
	GetByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByPhone(phone string) (*User, error)
}
