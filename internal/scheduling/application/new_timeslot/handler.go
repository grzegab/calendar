package new_timeslot

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

func (h *Handler) Handle(ctx context.Context, command Command) error {
	// validate user (if can create slots - is a teacher)
	// validate time (starts and ends in the future)
	scheduleTime, err := domain.NewScheduleTime(command.StartTime, command.EndTime)
	if err != nil {
		return err
	}

	schedule := domain.NewSchedule(command.TeacherID, scheduleTime)
	err = h.repo.Save(ctx, schedule)
	if err != nil {
		return err
	}

	return nil
}
