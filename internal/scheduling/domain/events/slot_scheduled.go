package events

import (
	"github/grzegab/calendar/internal/scheduling/domain"
	"time"
)

type SlotScheduled struct {
	SlotID    domain.SlotID
	TeacherID domain.TeacherID
}

func (e SlotScheduled) EventName() string {
	return "SlotScheduled"
}

func (e SlotScheduled) OccurredAt() time.Time {
	return time.Now()
}
