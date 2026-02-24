package decline_schedule

import (
	"context"
	"github/grzegab/calendar/internal/scheduling/domain"
)

type Handler struct {
	repo domain.SchedulingRepository
}

func NewHandler(repo domain.SchedulingRepository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context, id string) error {
	return nil
}
