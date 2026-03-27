package activate_user

import (
	"context"
	"github/grzegab/calendar/internal/shared/infrastructure/event_bus"
	"github/grzegab/calendar/internal/users/domain"
	"github/grzegab/calendar/internal/users/events"
)

type Handler struct {
	repo     domain.Repository
	eventBus event_bus.Bus
}

func NewHandler(repo domain.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context, id string) error {
	// do repo stuff, save active = 1 to user if exists

	h.eventBus.Publish(events.UserActivatedEvent{UserID: id})

	return nil
}
