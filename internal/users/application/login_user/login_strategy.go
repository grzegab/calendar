package login_user

import (
	"context"
	"errors"
)

type LoginStrategy interface {
	Authenticate(ctx context.Context, input LoginInput) (bool, error)
	Supports(input LoginInput) bool
}

type LoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginService struct {
	strategies []LoginStrategy
}

func NewLoginService(strategies ...LoginStrategy) *LoginService {
	return &LoginService{strategies: strategies}
}

func (ls LoginService) List() []LoginStrategy {
	return ls.strategies
}

func (ls LoginService) SelectStrategy(input LoginInput) (LoginStrategy, error) {
	for _, s := range ls.List() {
		if s.Supports(input) {
			return s, nil
		}
	}
	return nil, errors.New("no login strategy found")
}
