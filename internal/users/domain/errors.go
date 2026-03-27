package domain

import (
	"errors"
	"fmt"
)

type ErrUserNotFound struct {
	name string
}

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user %s not found", e.name)
}

var (
	ErrorUserNotFound        = errors.New("user not found")
	ErrorUserAlreadyInactive = errors.New("user already inactive")
	ErrorUserAlreadyActive   = errors.New("user already active")
)
