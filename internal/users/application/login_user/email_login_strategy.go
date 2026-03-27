package login_user

import (
	"context"
	"regexp"
)

type EmailLoginStrategy struct {
}

func (e *EmailLoginStrategy) Authenticate(ctx context.Context, input LoginInput) (bool, error) {
	// send email and wait for confirmation

	return true, nil
}

func (e *EmailLoginStrategy) Supports(input LoginInput) bool {
	if input.Password != "" {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(input.Login)
}
