package login_user

import (
	"context"
	"regexp"
)

type PhoneLoginStrategy struct {
}

func (p *PhoneLoginStrategy) Authenticate(ctx context.Context, input LoginInput) (bool, error) {
	// send sms and wait for confirmation

	return true, nil
}

func (p *PhoneLoginStrategy) Supports(input LoginInput) bool {
	if input.Password != "" {
		return false
	}

	cleanRegex := regexp.MustCompile(`[^\d+]`)
	cleanedPhone := cleanRegex.ReplaceAllString(input.Login, "")

	phoneRegex := regexp.MustCompile(`^(\+48)?\d{9}$`)

	return phoneRegex.MatchString(cleanedPhone)
}
