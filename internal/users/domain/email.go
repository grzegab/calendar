package domain

import (
	"strings"
)

type Email struct {
	value string
}

type InvalidEmailError struct {
	Value string
}

func (e InvalidEmailError) Error() string {
	return "invalid email: " + e.Value
}

func NewEmail(v string) (Email, error) {
	if !strings.Contains(v, "@") {
		return Email{}, InvalidEmailError{Value: v}
	}
	return Email{value: strings.ToLower(v)}, nil
}

func (e Email) Value() string {
	return e.value
}
