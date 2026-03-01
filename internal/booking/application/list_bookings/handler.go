package list_bookings

import "context"

type Handler struct {
	repo ReadRepository
}

func NewHandler(repo ReadRepository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context) error {
	return nil
}
