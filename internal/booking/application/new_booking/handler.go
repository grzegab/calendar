package new_booking

import (
	"context"
	"github/grzegab/calendar/internal/booking/domain"
)

type Handler struct {
	repo domain.BookingRepository
}

func NewHandler(repo domain.BookingRepository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context) error {
	return nil
}
