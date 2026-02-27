package register_user

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
	e, err := domain.NewEmail(command.Email)
	if err != nil {
		return err
	}

	p, err := domain.NewPhoneNumber(command.Phone)
	if err != nil {
		return err
	}

	user := domain.NewUser(e, p)

	return h.repo.Save(ctx, user)
}
