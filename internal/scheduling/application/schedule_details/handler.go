package schedule_details

import "github/grzegab/calendar/internal/scheduling/domain"

type Handler struct {
	repo domain.SchedulingRepository
}

func NewHandler(repo domain.SchedulingRepository) *Handler {
	return &Handler{repo: repo}
}
