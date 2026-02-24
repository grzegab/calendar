package domain

import (
	"github.com/google/uuid"
)

type SlotID string
type TeacherID string
type SlotStatus string

const (
	Open   SlotStatus = "open"
	Closed SlotStatus = "closed"
)

type Schedule struct {
	id           SlotID
	teacherID    TeacherID
	scheduleTime ScheduleTime
	status       SlotStatus
	events       []SchedulingEvent
}

func NewSchedule(teacherID string, scheduleTime ScheduleTime) *Schedule {
	// validate dates (start < end), max 2 h slot
	// max 8h of all slots per day

	return &Schedule{
		id:           SlotID(uuid.NewString()),
		teacherID:    TeacherID(teacherID),
		scheduleTime: scheduleTime,
		status:       Open,
	}
}

// SelectBooking will choose one of available bookings for timeslot. It will make schedule
// closed (no more bookings for this available).
func (s *Schedule) SelectBooking() {
	s.status = Closed
}
