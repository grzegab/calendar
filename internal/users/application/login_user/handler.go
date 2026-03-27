package login_user

import (
	"context"
	"errors"
	"fmt"
	"github/grzegab/calendar/internal/users/domain"
	"github/grzegab/calendar/internal/users/infrastructure/jwt_generator"
)

type Handler struct {
	repo         domain.Repository
	loginService *LoginService
	jwtGenerator jwt_generator.JwtGenerator
}

func NewHandler(repo domain.Repository, ls *LoginService, jwtg jwt_generator.JwtGenerator) *Handler {
	return &Handler{
		repo:         repo,
		loginService: ls,
		jwtGenerator: jwtg,
	}
}

// Handle returns JWT or error
func (h *Handler) Handle(ctx context.Context, query Query) (string, error) {
	strategy, err := h.loginService.SelectStrategy(query.LoginData)
	if err != nil {
		return "", fmt.Errorf("strategy error: %w\n", err)
	}

	if strategy == nil {
		return "", fmt.Errorf("strategy not found")
	}

	userValid, err := strategy.Authenticate(ctx, query.LoginData)
	if err != nil {
		return "", fmt.Errorf("strategy auth error: %w\n", err)
	}

	if !userValid {
		return "", errors.New("invalid user")
	}

	claims := jwt_generator.Claims{
		UserID: query.LoginData.Login,
	}
	jwtToken, err := h.jwtGenerator.Generate(claims)
	if err != nil {
		return "", fmt.Errorf("jwt generator error: %w\n", err)
	}

	return jwtToken, nil
}

func (h *Handler) selectStrategy(command Query) (LoginStrategy, error) {
	for _, s := range h.loginService.List() {
		if s.Supports(command.LoginData) {
			return s, nil
		}
	}
	return nil, errors.New("no login strategy found")
}
