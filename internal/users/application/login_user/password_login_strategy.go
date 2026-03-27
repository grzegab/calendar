package login_user

import (
	"context"
)

type PasswordLoginStrategy struct {
}

func (p *PasswordLoginStrategy) Authenticate(ctx context.Context, input LoginInput) (bool, error) {
	// bcrypt.CompareHashAndPassword - compare password in database
	return true, nil
}

func (p *PasswordLoginStrategy) Supports(input LoginInput) bool {
	if input.Login != "" && input.Password != "" {
		return true
	}

	return false
}
