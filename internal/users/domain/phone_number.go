package domain

import (
	"strings"
)

type PhoneNumber struct {
	value string
}

type InvalidPhoneNumberError struct {
	Value string
}

func (e InvalidPhoneNumberError) Error() string {
	return "invalid phone number: " + e.Value
}

func NewPhoneNumber(v string) (PhoneNumber, error) {
	cleaned := strings.ReplaceAll(v, " ", "")
	length := len(cleaned)

	if length == 9 {
		return PhoneNumber{value: cleaned}, nil
	}

	if length == 12 && len(cleaned) > 0 && cleaned[0] == '+' {
		return PhoneNumber{value: cleaned}, nil
	}

	return PhoneNumber{}, InvalidPhoneNumberError{Value: v}
}

func (p PhoneNumber) Value() string {
	return p.value
}
