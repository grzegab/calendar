package login_user

import (
	"context"
	"github/grzegab/calendar/internal/users/domain"
)

type Handler struct {
	repo domain.Repository
}

func NewHandler(repo domain.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context, command Command) error {
	return nil
}
